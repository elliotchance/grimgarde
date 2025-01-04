package main

import (
	"math"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type World struct {
	*tview.Box
	player *Player
	Map    *Map

	// This is where the player is located on to the map. This position is in the
	// center of the player sprite (which is the left side of the belt).
	PlayerX, PlayerY int

	// DestX and DestY is where the player is moving towards. When a new
	// destination is set, we precalculate the Distance so that each FrameNumber
	// is proportionally correct to the distance of travel. Otherwise, the lower
	// resolution of screen would make the movement look jagged.
	DestX, DestY int
	IsMoving     bool
	Distance     float64
	FrameNumber  int

	// StartX and StartY is where the player was when the Dest was set. We need to
	// keep this value as PlayerX and PlayerY will change over multiple frames
	// during the journey and the correct new position of each frame.
	StartX, StartY int
}

func NewWorld(m *Map, player *Player) *World {
	box := tview.NewBox()
	w := &World{
		Box:    box,
		player: player,
		Map:    m,
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
	w.IsMoving = true
	w.StartX, w.StartY = w.PlayerX, w.PlayerY
	w.DestX, w.DestY = x, y
	w.Distance = distance(w.StartX, w.StartY, w.DestX, w.DestY)
	w.FrameNumber = 0
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

	if w.IsMoving {
		tview.Print(screen, "@", w.DestX-w.PlayerX+x+(width/2), w.DestY-w.PlayerY+y+(height/2), 1, tview.AlignLeft, tcell.ColorWhite)
	}

	w.player.Draw(screen, x+(width/2), y+(height/2))
}

func (w *World) Start(app *tview.Application) {
	go func() {
		perSecond := 8.0
		ticker := time.NewTicker(time.Duration(1000/perSecond) * time.Millisecond)
		done := make(chan bool)

		go func() {
			for {
				select {
				case <-done:
					return
				case <-ticker.C:
					if w.IsMoving {
						w.FrameNumber++

						// w.Distance is the total diagonal length to Dest. We calculate
						// `traveled` as the length of the diagonal based on movement speed
						// and how many frames have passed. From this ideal location we can
						// calculate the correct PlayerX and PlayerY.
						traveled := (w.player.MovementSpeed() / perSecond) * float64(w.FrameNumber)
						portion := traveled / w.Distance
						w.PlayerX = w.StartX + int(math.Round(float64(w.DestX-w.StartX)*portion))
						w.PlayerY = w.StartY + int(math.Round(float64(w.DestY-w.StartY)*portion))
						w.IsMoving = portion < 1

						app.Draw()
					}
				}
			}
		}()
	}()
}

func distance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x1-x2), 2) + math.Pow(float64(y1-y2), 2))
}
