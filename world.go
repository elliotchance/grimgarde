package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type World struct {
	*tview.Box
	player          *Player
	playerPath      *Path
	Map             *Map
	FramesPerSecond int

	// This is where the player is located on to the map. This position is in the
	// center of the player sprite (which is the left side of the belt).
	PlayerX, PlayerY int

	Monsters []*Monster
}

func NewWorld(m *Map, player *Player, fps int) *World {
	box := tview.NewBox()
	w := &World{
		Box:             box,
		player:          player,
		playerPath:      NewEmptyPath(),
		Map:             m,
		FramesPerSecond: fps,
	}

	box.SetMouseCapture(func(action tview.MouseAction, event *tcell.EventMouse) (tview.MouseAction, *tcell.EventMouse) {
		if action == tview.MouseLeftClick {
			// (x, y) is the location in the screen. We have to correct for the
			// viewport and player location.
			x, y := event.Position()
			vx, vy, vw, vh := box.GetInnerRect()
			w.SetDest(w.PlayerX-(vw/2)+x-vx, w.PlayerY-(vh/2)+y+vy)
		}
		return action, event
	})

	return w
}

func (w *World) SetDest(x, y int) {
	w.playerPath = NewPath(w.PlayerX, w.PlayerY, x, y, w.player.MovementSpeed(), w.FramesPerSecond)
}

func (w *World) Viewport(width, height int) *Map {
	v := NewEmptyMap(width, height)
	x, y := w.PlayerX-width/2, w.PlayerY-height/2
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

	if w.playerPath.IsMoving {
		tview.Print(screen, "@",
			w.playerPath.DestX-w.PlayerX+x+(width/2), w.playerPath.DestY-w.PlayerY+y+(height/2),
			1, tview.AlignLeft, tcell.ColorWhite)
	}

	for _, monster := range w.Monsters {
		w.player.Draw(screen, monster.X-w.PlayerX+x+(width/2), monster.Y-w.PlayerY+y+(height/2))
	}

	w.player.Draw(screen, x+(width/2), y+(height/2))
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
					if w.playerPath.IsMoving {
						w.PlayerX, w.PlayerY = w.playerPath.Tick()
						redraw = true

						// When the player moves, the monsters should follow.
						for _, monster := range w.Monsters {
							monster.MoveTo(w.PlayerX, w.PlayerY)
						}
					}

					for _, monster := range w.Monsters {
						if monster.Path.IsMoving {
							redraw = true
							monster.Tick()
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
