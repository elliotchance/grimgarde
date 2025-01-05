package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type World struct {
	*tview.Box
	player          *Player
	Map             *Map
	FramesPerSecond int
	Monsters        []*Monster
}

func NewWorld(m *Map, player *Player, fps int) *World {
	box := tview.NewBox()
	w := &World{
		Box:             box,
		player:          player,
		Map:             m,
		FramesPerSecond: fps,
	}

	box.SetMouseCapture(func(action tview.MouseAction, event *tcell.EventMouse) (tview.MouseAction, *tcell.EventMouse) {
		if action == tview.MouseLeftClick {
			// (x, y) is the location in the screen. We have to correct for the
			// viewport and player location.
			x, y := event.Position()
			vx, vy, vw, vh := box.GetInnerRect()
			w.player.MoveTo(w.player.X-(vw/2)+x-vx, w.player.Y-(vh/2)+y+vy)
		}
		return action, event
	})

	return w
}

func (w *World) Viewport(width, height int) *Map {
	v := NewEmptyMap(width, height)
	x, y := w.player.X-width/2, w.player.Y-height/2
	for a := 0; a < height; a++ {
		for b := 0; b < width; b++ {
			if a+y >= 0 && a+y < len(w.Map.Data) &&
				b+x >= 0 && b+x < len(w.Map.Data[a+y]) {
				v.Data[a][b] = w.Map.Data[a+y][b+x]
			}
		}
	}
	return v
}

func (w *World) Draw(screen tcell.Screen) {
	x, y, width, height := w.Box.GetInnerRect()

	v := w.Viewport(width, height)
	for a := 0; a < width; a++ {
		for b := 0; b < height; b++ {
			ch := v.Data[b][a]
			tview.Print(screen, string(ch), x+a, y+b, 1, tview.AlignLeft, tcell.ColorWhite)
		}
	}

	if w.player.Path.IsMoving {
		w.Print(screen, w.player.Path.DestX, w.player.Path.DestY, "@")
	}

	for _, monster := range w.Monsters {
		w.player.Draw(screen, monster.X-w.player.X+x+(width/2), monster.Y-w.player.Y+y+(height/2))

		b := monster.Box(0, 0)
		w.Print(screen, b.x1, b.y1, "+")
		w.Print(screen, b.x2, b.y1, "+")
		w.Print(screen, b.x1, b.y2, "+")
		w.Print(screen, b.x2, b.y2, "+")
	}

	w.player.Draw(screen, x+(width/2), y+(height/2))
}

func (w *World) Print(screen tcell.Screen, atX, atY int, s string) {
	x, y, width, height := w.Box.GetInnerRect()
	tview.Print(screen, s,
		atX-w.player.X+x+(width/2), atY-w.player.Y+y+(height/2),
		len(s), tview.AlignLeft, tcell.ColorWhite)
}

func (w *World) Start(app *tview.Application) {
	go func() {
		ticker := time.NewTicker(time.Duration(1000/w.FramesPerSecond) * time.Millisecond)
		done := make(chan bool)

		go func() {
			for {
				select {
				case <-done:
					return
				case <-ticker.C:
					redraw := false
					if w.player.Path.IsMoving {
						w.player.X, w.player.Y = w.player.Path.Tick(func(b Box) bool {
							return true
						})
						redraw = true

						// When the player moves, the monsters should follow.
						for _, monster := range w.Monsters {
							monster.MoveTo(w.player.X, w.player.Y)
						}
					}

					for _, monster := range w.Monsters {
						if monster.Path.IsMoving {
							redraw = true
							monster.Tick(func(b Box) bool {
								return !b.Intersect(w.player.Box(w.player.X, w.player.Y))
							})
						}
					}

					if redraw {
						app.Draw()
					}
				}
			}
		}()
	}()
}
