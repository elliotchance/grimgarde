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
	m.Path = NewPath(m.X, m.Y, x, y, m.MovementSpeed, m.FramesPerSecond)
}

func (m *Monster) Tick() {
	m.X, m.Y = m.Path.Tick()
}
