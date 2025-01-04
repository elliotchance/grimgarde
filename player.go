package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Player struct {
	Level    int
	LevelExp int
	TotalExp int

	Strength     int // +Physical Damage +Physical Resistance
	Intelligence int // +Elemental Damage +Mana
	Dexterity    int // +Critical Chance +Dodge Chance
	Vitality     int // +Attack Speed +Life
	Unspent      int

	Life int
	Mana int

	Helmet *Item
	Chest  *Item
	Arms   *Item
	Gloves *Item
	Belt   *Item
	Pants  *Item
	Boots  *Item

	LeftHand  *Item
	RightHand *Item

	Amulet    *Item
	LeftRing  *Item
	RightRing *Item

	Inventory []*Item
}

func NewPlayer() *Player {
	return &Player{
		Level: 1,
	}
}

func (p *Player) AddExp(exp int) {
	p.TotalExp += exp
	nextLevel := p.NextLevelExp(p.Level)
	if p.LevelExp+exp >= nextLevel {
		p.AddLevel()
		exp -= nextLevel
	}
	p.LevelExp += exp
}

func (p *Player) AddLevel() {
	p.Level++
	p.Unspent += 5
	p.Fresh()
}

func (p *Player) Fresh() {
	p.Life = p.TotalLife()
	p.Mana = p.TotalMana()
}

// LevelExp returns how much exp is needed to get from level to level+1.
func (p *Player) NextLevelExp(level int) int {
	return int(1000 * math.Pow(1.1, float64(level-1)))
}

// PlayerStats is the combined stats of the player and items. This will return
// the combined attributes (eg. strength), but these have already been applied
// to the existing stats.
func (p *Player) PlayerStats() ItemStats {
	stats := ItemStats{}.Append(
		ItemStat{100, Life},
		ItemStat{50, Mana},
		ItemStat{p.Strength, Strength},
		ItemStat{p.Intelligence, Intelligence},
		ItemStat{p.Dexterity, Dexterity},
		ItemStat{p.Vitality, Vitality},
	)

	if p.Helmet != nil {
		stats = stats.Append(p.Helmet.Stats...)
	}

	if p.Chest != nil {
		stats = stats.Append(p.Chest.Stats...)
	}

	if p.Arms != nil {
		stats = stats.Append(p.Arms.Stats...)
	}

	if p.Gloves != nil {
		stats = stats.Append(p.Gloves.Stats...)
	}

	if p.Belt != nil {
		stats = stats.Append(p.Belt.Stats...)
	}

	if p.Pants != nil {
		stats = stats.Append(p.Pants.Stats...)
	}

	if p.Boots != nil {
		stats = stats.Append(p.Boots.Stats...)
	}

	if p.LeftHand != nil {
		// Scale weapon damage by attack rating
		scaled := p.LeftHand.Stats.Scale(p.LeftHand.AttackRating.AttacksPerSecond,
			PhysicalDamage, ColdDamage, FireDamage, PoisonDamage, LightningDamage)

		stats = stats.Append(scaled...)
	}

	if p.RightHand != nil {
		// Scale weapon damage by attack rating
		scaled := p.RightHand.Stats.Scale(p.RightHand.AttackRating.AttacksPerSecond,
			PhysicalDamage, ColdDamage, FireDamage, PoisonDamage, LightningDamage)

		stats = stats.Append(scaled...)
	}

	if p.Amulet != nil {
		stats = stats.Append(p.Amulet.Stats...)
	}

	if p.LeftRing != nil {
		stats = stats.Append(p.LeftRing.Stats...)
	}

	if p.RightRing != nil {
		stats = stats.Append(p.RightRing.Stats...)
	}

	// Apply attributes to stats
	stats = stats.Append(
		ItemStat{stats.Get(AllStats), Strength},
		ItemStat{stats.Get(AllStats), Intelligence},
		ItemStat{stats.Get(AllStats), Dexterity},
		ItemStat{stats.Get(AllStats), Vitality},
	)
	stats = stats.Append(
		ItemStat{stats.Get(Strength), PhysicalDamage},
		ItemStat{stats.Get(Strength), PhysicalResistance},
		ItemStat{stats.Get(Intelligence), ElementalDamage},
		ItemStat{stats.Get(Intelligence), Mana},
		ItemStat{stats.Get(Dexterity), CriticalChance},
		ItemStat{stats.Get(Dexterity), DodgeChance},
		ItemStat{stats.Get(Vitality), AttackSpeed},
		ItemStat{stats.Get(Vitality), Life},
	)

	return stats
}

func (p *Player) GetStat(statType ItemStatType) int {
	return p.PlayerStats().Get(statType)
}

func (p *Player) DamagePerSecond() int {
	stats := p.PlayerStats()
	damage := float64(stats.Get(PhysicalDamage) + stats.Get(ColdDamage) +
		stats.Get(FireDamage) + stats.Get(PoisonDamage) + stats.Get(LightningDamage))
	damage *= float64(stats.Get(AttackSpeed)) / 100
	return int(damage)
}

func (p *Player) TotalStrength() int {
	return p.GetStat(Strength) + p.GetStat(AllStats)
}

func (p *Player) TotalIntelligence() int {
	return p.GetStat(Intelligence) + p.GetStat(AllStats)
}

func (p *Player) TotalDexterity() int {
	return p.GetStat(Dexterity) + p.GetStat(AllStats)
}

func (p *Player) TotalVitality() int {
	return p.GetStat(Vitality) + p.GetStat(AllStats)
}

func (p *Player) TotalLife() int {
	return p.PlayerStats().Get(Life)
}

func (p *Player) TotalMana() int {
	return p.PlayerStats().Get(Mana)
}

func (p *Player) MovementSpeed() float64 {
	return 8
}

func (p *Player) String() string {
	items := []string{
		fmt.Sprintf("Level: %d", p.Level),
		fmt.Sprintf("Level Exp: %d of %d", p.LevelExp, p.NextLevelExp(p.Level)),
		fmt.Sprintf("Total Exp: %d", p.TotalExp),
		fmt.Sprintf("Strength: %d", p.Strength),
		fmt.Sprintf("Intelligence: %d", p.Intelligence),
		fmt.Sprintf("Dexterity: %d", p.Dexterity),
		fmt.Sprintf("Vitality: %d", p.Vitality),
		fmt.Sprintf("Unspent: %d", p.Unspent),
		fmt.Sprintf("Player Stats: %s", p.PlayerStats()),
		fmt.Sprintf("DPS: %d", p.DamagePerSecond()),
		fmt.Sprintf("Life: %d/%d", p.Life, p.TotalLife()),
		fmt.Sprintf("Mana: %d/%d", p.Mana, p.TotalMana()),
		p.Helmet.String(),
		p.Chest.String(),
		p.Arms.String(),
		p.Gloves.String(),
		p.Belt.String(),
		p.Pants.String(),
		p.Boots.String(),
		p.LeftHand.String(),
		p.RightHand.String(),
		p.Amulet.String(),
		p.LeftRing.String(),
		p.RightRing.String(),
	}

	return strings.Join(items, "\n\n")
}

func (p *Player) Draw(screen tcell.Screen, playerX, playerY int) {
	/*
			   ,,
		 *  '()'  /  )   /)
		  \ /||\ /  (   /
		   ' -- '
		    /  \
		   ^    ^
	*/

	draw := func(s string, x, y int) {
		tview.Print(screen, s, playerX+x, playerY+y, len(s), tview.AlignLeft, tcell.ColorWhite)
	}

	// Basic body
	draw(" () ", -1, -2)
	draw(" /||\\ ", -2, -1)
	draw(" ' -- ' ", -3, 0)
	draw(" /  \\ ", -2, 1)
	draw(" ^ ", -3, 2)
	draw(" ^ ", 2, 2)

	// Left mace
	draw(" * ", -5, -2)
	draw(" \\ ", -4, -1)

	// Right sword
	draw(" / ", 4, -2)
	draw(" / ", 3, -1)
}
