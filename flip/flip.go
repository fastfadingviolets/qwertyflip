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

const (
	ROW_SIZE = 10
	ROWS     = 4
)

func NewFlipper() Flipper {
	flipper := Flipper{
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

func (f *Flipper) RunCommand(command, line string) (string, error) {
	components := strings.Split(command, ",")
	for _, component := range components {
		if component == "H" {
			line = f.HorizontalFlip(line)
		} else if component == "V" {
			line = f.VerticalFlip(line)
		} else if places, err := strconv.Atoi(component); err == nil {
			line = f.Shift(places, line)
		} else {
			return "", fmt.Errorf("%s is an invalid command", component)
		}
	}
	return line, nil
}
