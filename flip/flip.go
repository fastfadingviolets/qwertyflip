package flip

import (
	"fmt"
	"strconv"
	"strings"
)

// Flipper has functions to qwerty-flip strings. Note that it isn't thread safe.
type Flipper struct {
	keyboard string
	keyMap   map[rune]int
	builder  strings.Builder
}

func NewFlipper() *Flipper {
	flipper := &Flipper{
		keyboard: "1234567890qwertyuiopasdfghjkl;zxcvbnm,./",
		keyMap:   make(map[rune]int),
		builder:  strings.Builder{},
	}
	for num, char := range flipper.keyboard {
		flipper.keyMap[char] = num
	}
	return flipper
}

func (f *Flipper) HorizontalFlip(line string) string {
	f.builder.Grow(len(line))
	defer f.builder.Reset()
	for _, char := range line {
		if pos, ok := f.keyMap[char]; ok {
			col := pos % 10
			newPos := pos - col + 9 - col
			f.builder.WriteByte(f.keyboard[newPos])
		} else {
			f.builder.WriteRune(char)
		}
	}
	return f.builder.String()
}

func (f *Flipper) VerticalFlip(line string) string {
	f.builder.Grow(len(line))
	defer f.builder.Reset()
	for _, char := range line {
		if pos, ok := f.keyMap[char]; ok {
			row := pos - (pos % 10)
			newPos := pos - row + 30 - row
			f.builder.WriteByte(f.keyboard[newPos])
		} else {
			f.builder.WriteRune(char)
		}
	}
	return f.builder.String()
}

func (f *Flipper) Shift(places int, line string) string {
	f.builder.Grow(len(line))
	defer f.builder.Reset()
	for _, char := range line {
		if pos, ok := f.keyMap[char]; ok {
			newPos := (pos + places) % 40
			if newPos < 0 {
				newPos = 40 + newPos
			}
			f.builder.WriteByte(f.keyboard[newPos])
		} else {
			f.builder.WriteRune(char)
		}
	}
	return f.builder.String()
}

type flipperFunction func(string) string

type Transform []flipperFunction

func (f *Flipper) ParseCommand(command string) (Transform, error) {
	parts := strings.Split(command, ",")
	retval := make([]flipperFunction, len(parts))
	for i, part := range parts {
		if part == "H" {
			retval[i] = f.HorizontalFlip
		} else if part == "V" {
			retval[i] = f.VerticalFlip
		} else if places, err := strconv.Atoi(part); err == nil {
			retval[i] = func(line string) string { return f.Shift(places, line) }
		} else if part == "" {
			// just in case there's a trailing comma or similar
			retval[i] = func(line string) string { return line }
		} else {
			return nil, fmt.Errorf("%s is an invalid command", part)
		}
	}
	return retval, nil
}

func (t Transform) Apply(line string) string {
	for _, funct := range t {
		line = funct(line)
	}
	return line
}
