package main

type World struct {
	Map *Map
}

func NewWorld(m *Map) *World {
	return &World{
		Map: m,
	}
}

func (e *World) Viewport(x, y, h, w int) *Map {
	v := NewEmptyMap(h, w)
	for a := 0; a < h; a++ {
		for b := 0; b < w; b++ {
			v.Data[a][b] = e.Map.Data[a+y][b+x]
		}
	}
	return v
}
