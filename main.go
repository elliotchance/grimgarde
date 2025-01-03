package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

/*
     ,,
 *  '()'  /  )   /)
  \ /||\ /  (   /
   ' -- '
    /  \
   ^    ^

*/

func main() {
	// p := NewPlayer()
	// p.AddExp(1500)
	// p.AddExp(200)
	// p.AddExp(500)

	// p.Strength += 3
	// p.Intelligence += 5
	// p.Dexterity += 7
	// p.Vitality += 11

	// p.Helmet = GenerateHelmet()
	// p.Chest = GenerateChest()
	// p.Arms = GenerateArms()
	// p.Gloves = GenerateGloves()
	// p.Belt = GenerateBelt()
	// p.Pants = GeneratePants()
	// p.Boots = GenerateBoots()
	// p.LeftHand = GenerateWeapon()
	// p.RightHand = GenerateWeapon()
	// p.Amulet = GenerateAmulet()
	// p.LeftRing = GenerateRing()
	// p.RightRing = GenerateRing()

	// p.Fresh()

	// fmt.Println(p.String())

	playerX := 4
	playerY := 4

	world := NewWorld(townMap)
	// v := world.Viewport(playerX-4, playerY-4, 9, 9)
	// fmt.Println(v)
	// panic("")

	box := tview.NewBox().SetBorder(true).SetTitle("Character")
	canvas := tview.NewBox()
	canvas.SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		x, y, w, h := canvas.GetInnerRect()

		v := world.Viewport(playerX-4, playerY-4, 9, 9)
		// fmt.Println(v)

		for a := 0; a < v.Width; a++ {
			for b := 0; b < v.Height; b++ {
				ch := v.Data[b][a]
				// color := "black"
				// if ch == '0' || ch == 'd' {
				// 	color = "red"
				// }
				tview.Print(screen, string(ch), x+a, y+b, 1, tview.AlignLeft, tcell.ColorWhite)
			}
		}

		tview.Print(screen, "🧍", x+4, y+4, 1, tview.AlignLeft, tcell.ColorWhite)
		return x, y, w, h
	})
	// status := tview.NewTextView().SetText("status bar")
	grid := tview.NewGrid().AddItem(box, 0, 0, 1, 1, 1, 1, false).
		AddItem(canvas, 0, 1, 1, 1, 1, 1, true)
		// AddItem(status, 1, 0, 1, 2, 1, 1, false)
	if err := tview.NewApplication().
		SetRoot(grid, true).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case tcell.KeyLeft:
				playerX--
			case tcell.KeyRight:
				playerX++
			case tcell.KeyUp:
				playerY--
			case tcell.KeyDown:
				playerY++
			}
			if playerX < 4 {
				playerX = 4
			}
			if playerY < 4 {
				playerY = 4
			}
			return event
		}).Run(); err != nil {
		panic(err)
	}
}
