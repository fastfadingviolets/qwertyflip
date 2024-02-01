package main

import (
	"strings"
)

type rowColumn struct {
	Row    int
	Column int
}

// Flipper has functions to qwerty-flip strings. Note that it isn't thread safe.
type Flipper struct {
	keyboard [4]string
	keyMap   map[rune]rowColumn
	builder  strings.Builder
}

func NewFlipper() Flipper {
	flipper := Flipper{
		keyboard: [...]string{"1234567890", "qwertyuiop", "asdfghjkl;", "zxcvbnm,./"},
		keyMap:   make(map[rune]rowColumn),
		builder:  strings.Builder{},
	}
	for rowNum, row := range flipper.keyboard {
		for colNum, letter := range row {
			flipper.keyMap[letter] = rowColumn{rowNum, colNum}
		}
	}
	return flipper
}

func (f *Flipper) HorizontalFlip(line string) string {
	f.builder.Grow(len(line))
	defer f.builder.Reset()
	for _, char := range line {
		if rowCol, ok := f.keyMap[char]; ok {
			newCol := 9 - rowCol.Column
			newChar := f.keyboard[rowCol.Row][newCol]
			f.builder.WriteByte(newChar)
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
		if rowCol, ok := f.keyMap[char]; ok {
			newRow := 3 - rowCol.Row
			newChar := f.keyboard[newRow][rowCol.Column]
			f.builder.WriteByte(newChar)
		} else {
			f.builder.WriteRune(char)
		}
	}
	return f.builder.String()
}

func (f *Flipper) Shift(places int, line string) string {
	f.builder.Grow(len(line))
	defer f.builder.Reset()
	return f.builder.String()
}
