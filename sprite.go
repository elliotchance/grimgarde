package main

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Sprite struct {
	Down          []string
	Width, Height int
}

func NewSprite(width, height int) *Sprite {
	return &Sprite{
		Width:  width,
		Height: height,
	}
}

func (s *Sprite) SetDown(sp string) *Sprite {
	s.Down = s.normalizeSprite(sp)
	return s
}

func (s *Sprite) Draw(screen tcell.Screen, x, y int) {
	// The (x, y) refers to the center.
	x1 := x - (s.Width / 2)
	y1 := y - (s.Height / 2)
	for i, line := range s.Down {
		tview.Print(screen, line, x1, y1+i, len(line), tview.AlignLeft, tcell.ColorWhite)
	}
}

func (s *Sprite) Box(x, y int) Box {
	// The (x, y) refers to the center.
	x1 := x - (s.Width / 2)
	y1 := y - (s.Height / 2)
	return NewBox(x1, y1, x1+s.Width-1, y1+s.Height-1)
}

func (s *Sprite) normalizeSprite(data string) []string {
	sprite := strings.Split(strings.Trim(string(data), "\n"), "\n")

	// Check/fix widths
	if len(sprite) != s.Height {
		panic(fmt.Sprintf("sprite wrong height (%d != %d)", len(sprite), s.Height))
	}
	for i, line := range sprite {
		if len(line) > s.Width {
			panic("sprite too wide")
		}
		if len(line) < s.Width {
			sprite[i] += strings.Repeat(" ", s.Width-len(line))
		}
	}

	return sprite
}
