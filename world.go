package main

type World struct {
	Map  *Map
	X, Y int
}

func NewWorld(m *Map) *World {
	return &World{
		Map: m,
	}
}

func (w *World) Viewport(width, height int) *Map {
	v := NewEmptyMap(width, height)
	for a := 0; a < height; a++ {
		for b := 0; b < width; b++ {
			v.Data[a][b] = w.Map.Data[a+w.Y][b+w.X]
		}
	}
	return v
}
