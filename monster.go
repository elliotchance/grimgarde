package main

type Monster struct {
	Type            string
	Life, MaxLife   int
	MovementSpeed   float64
	X, Y            int
	Path            *Path
	FramesPerSecond int
}

func NewSpider(x, y, fps int) *Monster {
	return &Monster{
		Type:            "Spider",
		Life:            100,
		MaxLife:         100,
		MovementSpeed:   4,
		X:               x,
		Y:               y,
		FramesPerSecond: fps,
		Path:            NewEmptyPath(),
	}
}

func (m *Monster) MoveTo(x, y int) {
	m.Path = NewPath(m.X, m.Y, x, y, m.MovementSpeed, m.FramesPerSecond, m)
}

func (m *Monster) Tick(canMove func(b Box) bool) {
	m.X, m.Y = m.Path.Tick(canMove)
}

func (m *Monster) Box(x, y int) Box {
	// The (x, y) refers to the center.
	return NewBox(x-4, y-2, x+5, y+2)
}
