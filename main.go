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

	for i := 0; i < 10; i++ {
		player.Inventory = append(player.Inventory, GenerateAnyItem())
	}

	player.Fresh()

	app := tview.NewApplication()
	framesPerSecond := 8
	world := NewWorld(townMap, player, framesPerSecond)
	world.PlayerX, world.PlayerY = 75, 25

	spider := NewSpider(45, 28, framesPerSecond)
	spider.MoveTo(world.PlayerX, world.PlayerY)
	world.Monsters = append(world.Monsters, spider)

	world.Start(app)

	grid := tview.NewFlex().
		AddItem(world, 0, 1, true)

	outerGrid := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(grid, 0, 1, false).
		AddItem(NewStatusView(player), 1, 1, false)

	characterScreenIsOpen := false
	inventoryScreenIsOpen := false

	if err := app.
		SetRoot(outerGrid, true).
		EnableMouse(true).
		SetFocus(grid).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyRune {
				grid.Clear()
				switch event.Rune() {
				case 'c':
					if characterScreenIsOpen {
						grid.AddItem(world, 0, 1, false)
						app.SetFocus(world)
					} else {
						characterView := NewCharacterView(player)
						grid.AddItem(characterView, 70, 1, false)
						grid.AddItem(world, 0, 1, false)
						app.SetFocus(characterView)
					}
					characterScreenIsOpen = !characterScreenIsOpen
				case 'i':
					if inventoryScreenIsOpen {
						grid.AddItem(world, 0, 1, false)
						app.SetFocus(world)
					} else {
						inventory := NewInventoryView(player)
						grid.AddItem(inventory, 70, 1, false).
							AddItem(world, 0, 1, false)
						app.SetFocus(inventory)
					}
					inventoryScreenIsOpen = !inventoryScreenIsOpen
				}

				return event
			}

			return event
		}).
		Run(); err != nil {
		panic(err)
	}
}
