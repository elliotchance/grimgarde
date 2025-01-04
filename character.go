package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type CharacterView struct {
	*tview.Form
	player *Player
}

func NewCharacterView(player *Player) *CharacterView {
	form := tview.NewForm()
	form.SetTitle("Character")
	form.SetBorder(true)

	return &CharacterView{
		Form:   form,
		player: player,
	}
}

func (g *CharacterView) Draw(screen tcell.Screen) {
	g.Form.Clear(true)

	g.Form.
		AddInputField("Name", "", 20, nil, nil).
		AddTextView("Level", fmt.Sprintf("%d (%d of %d)", g.player.Level, g.player.LevelExp, g.player.NextLevelExp(g.player.Level)), 20, 1, false, false).
		AddTextView("Life", fmt.Sprintf("%d / %d", g.player.Life, g.player.TotalLife()), 10, 1, false, false).
		AddTextView("Mana", fmt.Sprintf("%d / %d", g.player.Mana, g.player.TotalMana()), 10, 1, false, false).
		AddTextView("Strength", fmt.Sprintf("%d", g.player.TotalStrength()), 10, 1, false, false).
		AddTextView("Intelligence", fmt.Sprintf("%d", g.player.TotalIntelligence()), 10, 1, false, false).
		AddTextView("Dexterity", fmt.Sprintf("%d", g.player.TotalDexterity()), 10, 1, false, false).
		AddTextView("Vitality", fmt.Sprintf("%d", g.player.TotalVitality()), 10, 1, false, false).
		AddTextView("Unspent", fmt.Sprintf("%d", g.player.Unspent), 10, 1, false, false)
	if g.player.Unspent > 0 {
		g.Form.AddButton("+ Strength", func() {
			g.player.Strength++
			g.player.Unspent--
		})
		g.Form.AddButton("+ Intelligence", func() {
			g.player.Intelligence++
			g.player.Unspent--
		})
		g.Form.AddButton("+ Dexterity", func() {
			g.player.Dexterity++
			g.player.Unspent--
		})
		g.Form.AddButton("+ Vitality", func() {
			g.player.Vitality++
			g.player.Unspent--
		})
	}

	g.Form.Draw(screen)
}
