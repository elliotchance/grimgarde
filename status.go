package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type StatusView struct {
	*tview.Flex
	life   *tview.TextView
	mana   *tview.TextView
	width  int
	player *Player
}

func NewStatusView(player *Player) *StatusView {
	life := tview.NewTextView().SetDynamicColors(true)
	mana := tview.NewTextView().SetDynamicColors(true)
	width := 30

	flex := tview.NewFlex().
		AddItem(life, width, 1, false).
		AddItem(tview.NewTextView().SetText("   [C]haractor   [I]nventory"), 0, 1, false).
		AddItem(mana, width, 1, false)

	return &StatusView{
		Flex:   flex,
		life:   life,
		mana:   mana,
		width:  width,
		player: player,
	}
}

func (g *StatusView) Draw(screen tcell.Screen) {
	totalLife := g.player.TotalLife()
	lifePixels := int(math.Ceil((float64(g.player.Life) / float64(totalLife)) * float64(g.width)))
	s := fmt.Sprintf(" %d / %d", g.player.Life, totalLife)
	s += strings.Repeat(" ", g.width-len(s))
	s = s[:lifePixels] + "[:darkred]" + s[lifePixels:]
	s = "[:red]" + s
	g.life.SetText(s)

	totalMana := g.player.TotalMana()
	manaPixels := g.width - int(math.Ceil((float64(g.player.Mana)/float64(totalMana))*float64(g.width)))
	s = fmt.Sprintf("%d / %d ", g.player.Mana, totalMana)
	s = strings.Repeat(" ", g.width-len(s)) + s
	s = s[:manaPixels] + "[:blue]" + s[manaPixels:]
	s = "[:darkblue]" + s
	g.mana.SetText(s)

	g.Flex.Draw(screen)
}
