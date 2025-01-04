package main

import (
	"sort"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type InventoryView struct {
	*tview.Flex
	inventory    *tview.Table
	selectedItem *tview.Table
	equippedItem *tview.Table
	player       *Player
}

func NewInventoryView(player *Player) *InventoryView {
	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)
	flex.SetTitle("Inventory")
	flex.SetBorder(true)

	inventory := tview.NewTable()
	inventory.SetSelectable(true, false)
	inventory.SetEvaluateAllRows(true)
	flex.AddItem(inventory, 0, 1, true)

	selectedItem := tview.NewTable()
	selectedItem.SetBorder(true)

	equippedItem := tview.NewTable()
	equippedItem.SetBorder(true)

	compare := tview.NewGrid()
	compare.AddItem(selectedItem, 0, 0, 1, 1, 0, 0, true)
	compare.AddItem(equippedItem, 0, 1, 1, 1, 0, 0, false)

	// 11 is the max lines for an item.
	flex.AddItem(compare, 11, 1, false)

	return &InventoryView{
		Flex:         flex,
		inventory:    inventory,
		selectedItem: selectedItem,
		equippedItem: equippedItem,
		player:       player,
	}
}

func (g *InventoryView) Draw(screen tcell.Screen) {
	g.inventory.SetCell(0, 0, tview.NewTableCell("* Helmet"))
	g.inventory.SetCell(0, 1, tview.NewTableCell(g.player.Helmet.Name))
	g.inventory.SetCell(1, 0, tview.NewTableCell("* Chest"))
	g.inventory.SetCell(1, 1, tview.NewTableCell(g.player.Chest.Name))
	g.inventory.SetCell(2, 0, tview.NewTableCell("* Arms"))
	g.inventory.SetCell(2, 1, tview.NewTableCell(g.player.Arms.Name))
	g.inventory.SetCell(3, 0, tview.NewTableCell("* Gloves"))
	g.inventory.SetCell(3, 1, tview.NewTableCell(g.player.Gloves.Name))
	g.inventory.SetCell(4, 0, tview.NewTableCell("* Pants"))
	g.inventory.SetCell(4, 1, tview.NewTableCell(g.player.Pants.Name))
	g.inventory.SetCell(5, 0, tview.NewTableCell("* Boots"))
	g.inventory.SetCell(5, 1, tview.NewTableCell(g.player.Boots.Name))
	g.inventory.SetCell(6, 0, tview.NewTableCell("* Weapon"))
	g.inventory.SetCell(6, 1, tview.NewTableCell(g.player.LeftHand.Name))
	g.inventory.SetCell(7, 0, tview.NewTableCell("* Weapon"))
	g.inventory.SetCell(7, 1, tview.NewTableCell(g.player.RightHand.Name))
	g.inventory.SetCell(8, 0, tview.NewTableCell("* Ring"))
	g.inventory.SetCell(8, 1, tview.NewTableCell(g.player.LeftRing.Name))
	g.inventory.SetCell(9, 0, tview.NewTableCell("* Ring"))
	g.inventory.SetCell(9, 1, tview.NewTableCell(g.player.RightRing.Name))
	g.inventory.SetCell(10, 0, tview.NewTableCell("* Amulet"))
	g.inventory.SetCell(10, 1, tview.NewTableCell(g.player.Amulet.Name))
	equippedRows := 10

	sort.Slice(g.player.Inventory, func(i, j int) bool {
		if g.player.Inventory[i].Type.Class != g.player.Inventory[j].Type.Class {
			return g.player.Inventory[i].Type.Class < g.player.Inventory[j].Type.Class
		}

		return g.player.Inventory[i].Name < g.player.Inventory[j].Name
	})

	for i, item := range g.player.Inventory {
		g.inventory.SetCell(equippedRows+i, 0, tview.NewTableCell(string(item.Type.Class)))
		g.inventory.SetCell(equippedRows+i, 1, tview.NewTableCell(item.Name))
	}

	selectedRow, _ := g.inventory.GetSelection()
	selectedText := g.inventory.GetCell(selectedRow, 0).Text
	var equippedItem *Item
	if strings.HasPrefix(selectedText, "* ") {
		equippedItem = g.getEquipped(selectedText[2:])
	} else {
		equippedItem = g.getEquipped(selectedText)
	}

	g.equippedItem.Clear()
	if equippedItem != nil {
		g.equippedItem.SetTitle("Equipped")
		for i, stat := range equippedItem.Lines() {
			g.equippedItem.SetCell(i, 0, tview.NewTableCell(stat))
		}
	} else {
		g.equippedItem.SetTitle("None Equipped")
	}

	g.selectedItem.Clear()
	g.selectedItem.SetTitle("Selected")
	if selectedRow >= equippedRows {
		for i, stat := range g.player.Inventory[selectedRow-equippedRows].Lines() {
			g.selectedItem.SetCell(i, 0, tview.NewTableCell(stat))
		}
	}

	g.Flex.Draw(screen)
}

func (g *InventoryView) getEquipped(className string) *Item {
	switch className {
	case "Helmet":
		return g.player.Helmet
	case "Chest":
		return g.player.Chest
	case "Arms":
		return g.player.Arms
	case "Gloves":
		return g.player.Gloves
	case "Pants":
		return g.player.Pants
	case "Boots":
		return g.player.Boots
	case "Weapon":
		return g.player.LeftHand
	case "Ring":
		return g.player.LeftRing
	case "Amulet":
		return g.player.Amulet
	}

	return nil
}

func (g *InventoryView) Focus(delegate func(p tview.Primitive)) {
	g.Flex.Focus(delegate)
}

func (g *InventoryView) HasFocus() bool {
	return g.Flex.HasFocus()
}
