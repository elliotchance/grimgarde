package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type Monster struct {
	Type            string
	Life, MaxLife   int
	MovementSpeed   float64
	X, Y            int
	Path            *Path
	FramesPerSecond int
	Damage          int
	LastAttack      time.Time

	// AttackSpeed is the number of attacks per second.
	AttackSpeed float64

	Sprite *Sprite
}

func NewSpider(x, y, fps int) *Monster {
	return &Monster{
		Type:            "Spider",
		Life:            100,
		MaxLife:         100,
		X:               x,
		Y:               y,
		FramesPerSecond: fps,
		Path:            NewEmptyPath(),
		AttackSpeed:     0.5,
		Damage:          10,

		// MovementSpeed must be a bit slower than the player so it's out runnable.
		MovementSpeed: NormalMovingSpeed * 0.5,

		Sprite: NewSprite(8, 4).SetDown(`
 ||  ||
 \\()//
//(__)\\
||    ||
`),
	}
}

func (m *Monster) MoveTo(x, y int) {
	m.Path = NewPath(m.X, m.Y, x, y, m.MovementSpeed, m.FramesPerSecond, m)
}

func (m *Monster) Tick(canMove func(b Box) bool) {
	m.X, m.Y = m.Path.Tick(canMove)
}

func (m *Monster) Box(x, y int) Box {
	return m.Sprite.Box(m.X, m.Y)
}

func (m *Monster) Attack(player *Player) bool {
	minTimeBetweenAttacks := time.Duration(float64(time.Second) / m.AttackSpeed)
	if time.Since(m.LastAttack) >= minTimeBetweenAttacks {
		m.LastAttack = time.Now()
		player.Hit(m.Damage)
		return true
	}

	return false
}

func (m *Monster) Draw(screen tcell.Screen, x, y int) {
	m.Sprite.Draw(screen, x, y)
}
