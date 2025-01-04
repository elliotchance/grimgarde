package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	player := NewPlayer()
	player.AddExp(1500)
	player.AddExp(200)
	player.AddExp(500)

	player.Strength += 3
	player.Intelligence += 5
	player.Dexterity += 7
	player.Vitality += 11

	player.Helmet = GenerateHelmet()
	player.Chest = GenerateChest()
	player.Arms = GenerateArms()
	player.Gloves = GenerateGloves()
	player.Belt = GenerateBelt()
	player.Pants = GeneratePants()
	player.Boots = GenerateBoots()
	player.LeftHand = GenerateWeapon()
	player.RightHand = GenerateWeapon()
	player.Amulet = GenerateAmulet()
	player.LeftRing = GenerateRing()
	player.RightRing = GenerateRing()

	player.Fresh()

	// fmt.Println(p.String())

	world := NewWorld(townMap)

	canvas := tview.NewBox()
	canvas.SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		x, y, w, h := canvas.GetInnerRect()

		v := world.Viewport(w, h)
		for a := 0; a < v.Width; a++ {
			for b := 0; b < v.Height; b++ {
				ch := v.Data[b][a]
				tview.Print(screen, string(ch), x+a, y+b, 1, tview.AlignLeft, tcell.ColorWhite)
			}
		}

		// Render player
		player.Draw(screen, x+(w/2), y+(h/2))

		return x, y, w, h
	})
	// status := tview.NewTextView().SetText("status bar")
	app := tview.NewApplication()
	grid := tview.NewGrid().AddItem(NewCharacterView(player), 0, 0, 1, 1, 1, 1, false).
		AddItem(canvas, 0, 1, 1, 1, 1, 1, true)
	// AddItem(status, 1, 0, 1, 2, 1, 1, false)
	if err := app.
		SetRoot(grid, true).
		EnableMouse(true).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case tcell.KeyLeft:
				world.X--
			case tcell.KeyRight:
				world.X++
			case tcell.KeyUp:
				world.Y--
			case tcell.KeyDown:
				world.Y++
			}
			if world.X < 0 {
				world.X = 0
			}
			if world.Y < 0 {
				world.Y = 0
			}
			return event
		}).
		Run(); err != nil {
		panic(err)
	}
}
