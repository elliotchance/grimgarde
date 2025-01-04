package main

import (
	"fmt"
	"os"
	"strings"
)

var sprites = map[rune][]string{
	// Punctuation is applied first.
	'.': loadSprite("grass"),
	'#': loadSprite("path"),
	'~': loadSprite("water"),
	'+': loadSprite("fence-corner"),
	'-': loadSprite("fence-horizontal"),
	'|': loadSprite("fence-vertical"),

	// Then background sprites.
	'h': loadSprite("house"),

	// Finally, foreground sprites.
	'N': loadSprite("npc"),
}

// Resolution is 5px per character
var townMap = NewMap(`
+----------------------------------------------------------+
|~~~~~~~~~~~~~~~~###.......................................|
|~~~~~~~~~~~~~~~~###.......................................|
|~~~~~~~~~~~~~~~~###h           ...........................|
|~~~~~~~~~~~~~~~~###            ...........................|
|~~~~~~~~~~~~~~~~###            ...........................|
|~~~~~~~~~~~~~~~~###            ...........................|
|~~~~~~~~~~~~~~~~###            ...........................|
|~~~~~~~~~~~~~~~~###            ...........................|
|~~~~~~~~~~~~~~~~###            ...........................|
|~~~~~~~~~~~~~~~~###            ...........................|
|~~~~~~~~~~~~~~~~###.N.....................................|
|~~~~~~~~~~~~~~~~###.......................................|
|~~~~~~~~~~~~~~~~###.......................................|
|~~~~~~~~~~~~~~~~###.......................................|
|~~~~~~~~~~~~~~~~###.......................................|
|~~~~~~~~~~~~~~~~###.......................................|
|~~~~~~~~~~~~~~~~###.......................................|
|~~~~~~~~~~~~~~~~###.......................................|
+----------------------------------------------------------+
`)

type Map struct {
	Data          [][]rune
	Width, Height int
}

func NewEmptyMap(width, height int) *Map {
	m := &Map{
		Width:  width,
		Height: height,
		Data:   make([][]rune, height),
	}
	for i := range m.Data {
		m.Data[i] = make([]rune, width)
	}

	return m
}

func NewMap(pixels string) *Map {
	pixels = strings.TrimSpace(pixels)
	rows := strings.Split(pixels, "\n")
	height := len(rows)
	if height == 0 {
		panic("empty map")
	}
	width := len(rows[0])
	for i, row := range rows {
		if len(row) != width {
			panic(fmt.Sprintf("incorrect length for row %d", i+1))
		}
	}

	// Render
	m := NewEmptyMap(width*5, height*5)
	for i, row := range rows {
		for offset, pixel := range row {
			if pixel == ' ' {
				continue
			}
			sprite := sprites[pixel]
			if len(sprite) == 0 {
				panic(fmt.Sprintf("unknown pixel: %c", pixel))
			}
			for j := range sprite {
				for k := range sprite[j] {
					m.Data[i*5+j][offset*5+k] = rune(sprite[j][k])
				}
			}
		}
	}

	return m
}

func (m *Map) String() string {
	var mapLines []string
	for _, row := range m.Data {
		mapLines = append(mapLines, string(row))
	}

	return strings.Join(mapLines, "\n")
}

func loadSprite(name string) []string {
	data, err := os.ReadFile("sprites/" + name + ".txt")
	if err != nil {
		panic(err)
	}

	sprite := strings.Split(strings.Trim(string(data), "\n"), "\n")
	if len(sprite)%5 != 0 || len(sprite) == 0 {
		panic(fmt.Sprintf("sprite %s: incorrect height", name))
	}

	// Pad out width, if needed.
	maxWidth := 0
	for _, line := range sprite {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	if maxWidth%5 != 0 {
		maxWidth += maxWidth % 5
	}

	for i, line := range sprite {
		if len(line) != maxWidth {
			sprite[i] += strings.Repeat(" ", maxWidth-len(sprite[i]))
		}
	}

	return sprite
}
