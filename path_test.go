package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type NoopBoxer struct{}

func (n NoopBoxer) Box(x, y int) Box {
	return NewBox(x, y, x+3, y+3)
}

func TestPath_Tick(t *testing.T) {
	p := NewPath(3, 3, 10, 12, 4, 2, NoopBoxer{})
	assertTick(t, p, 4, 5, true)    // 0.5s
	assertTick(t, p, 5, 6, true)    // 1s
	assertTick(t, p, 7, 8, true)    // 1.5s
	assertTick(t, p, 8, 9, true)    // 2s
	assertTick(t, p, 9, 11, true)   // 2.5s
	assertTick(t, p, 10, 12, false) // 3s
	assertTick(t, p, 10, 12, false) // 3.5s
}

func assertTick(t *testing.T, path *Path, expectedX, expectedY int, isMoving bool) {
	t.Helper()
	x, y := path.Tick(func(b Box) bool {
		return true
	})
	assert.Equal(t, expectedX, x)
	assert.Equal(t, expectedY, y)
	assert.Equal(t, isMoving, path.IsMoving)
}
