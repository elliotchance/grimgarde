package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBox_Touching(t *testing.T) {
	/*
			0 >

			2 >      +-+
			         |2|
		  4 >      +-+
		           ^ ^
			6 >  +-+ . .
			     |1| . .
			8 >  +-+ . .
		     ^ ^ ^ . .
			   0 2 4 6 8
	*/

	// b1 runs underneath b2 with a gap of 1
	for x := 0; x < 9; x++ {
		assert.False(t, NewBox(2+x, 6, 4+x, 8).Touching(NewBox(6, 2, 8, 4)), x)
	}
	// b1 runs underneath b2 with a gap of 0 (touching)
	assert.False(t, NewBox(2+0, 5, 4+0, 7).Touching(NewBox(6, 2, 8, 4)))
	assert.True(t, NewBox(2+1, 5, 4+1, 7).Touching(NewBox(6, 2, 8, 4)))
	assert.True(t, NewBox(2+2, 5, 4+2, 7).Touching(NewBox(6, 2, 8, 4)))
	assert.True(t, NewBox(2+3, 5, 4+3, 7).Touching(NewBox(6, 2, 8, 4)))
	assert.True(t, NewBox(2+4, 5, 4+4, 7).Touching(NewBox(6, 2, 8, 4)))
	assert.True(t, NewBox(2+5, 5, 4+5, 7).Touching(NewBox(6, 2, 8, 4)))
	assert.True(t, NewBox(2+6, 5, 4+6, 7).Touching(NewBox(6, 2, 8, 4)))
	assert.True(t, NewBox(2+7, 5, 4+7, 7).Touching(NewBox(6, 2, 8, 4)))
	assert.False(t, NewBox(2+8, 5, 4+8, 7).Touching(NewBox(6, 2, 8, 4)))

	// b1 runs to the left of b2 with a gap of 1
	for y := 0; y < 9; y++ {
		assert.False(t, NewBox(2, 6-y, 4, 8-y).Touching(NewBox(6, 2, 8, 4)), y)
	}
	// b1 runs to the left of b2 with a gap of 0 (touching)
	assert.False(t, NewBox(2, 5-0, 4, 7-0).Touching(NewBox(6, 2, 8, 4)))
	assert.False(t, NewBox(2, 5-1, 4, 7-1).Touching(NewBox(6, 2, 8, 4)))
	assert.False(t, NewBox(2, 5-2, 4, 7-2).Touching(NewBox(6, 2, 8, 4)))
	assert.False(t, NewBox(2, 5-3, 4, 7-3).Touching(NewBox(6, 2, 8, 4)))
	assert.False(t, NewBox(2, 5-4, 4, 7-4).Touching(NewBox(6, 2, 8, 4)))
	assert.False(t, NewBox(2, 5-5, 4, 7-5).Touching(NewBox(6, 2, 8, 4)))
	assert.False(t, NewBox(2, 5-6, 4, 7-6).Touching(NewBox(6, 2, 8, 4)))
	assert.False(t, NewBox(2, 5-7, 4, 7-7).Touching(NewBox(6, 2, 8, 4)))
	assert.False(t, NewBox(2, 5-8, 4, 7-8).Touching(NewBox(6, 2, 8, 4)))

	/*
			0 >

			2 >  +-+
			     |1|
		  4 >  +-+
		       ^ ^
			6 >  . . +-+
			     . . |2|
			8 >  . . +-+
		     ^ . . ^ ^
			   0 2 4 6 8
	*/

	// b1 runs above b2 with a gap of 1
	for x := 0; x < 9; x++ {
		assert.False(t, NewBox(2+x, 2, 4+x, 4).Touching(NewBox(6, 6, 8, 8)), x)
	}
	// b1 runs above b2 with a gap of 0 (touching)
	assert.False(t, NewBox(2+0, 2, 4+0, 4).Touching(NewBox(6, 5, 8, 7)))
	assert.True(t, NewBox(2+1, 2, 4+1, 4).Touching(NewBox(6, 5, 8, 7)))
	assert.True(t, NewBox(2+2, 2, 4+2, 4).Touching(NewBox(6, 5, 8, 7)))
	assert.True(t, NewBox(2+3, 2, 4+3, 4).Touching(NewBox(6, 5, 8, 7)))
	assert.True(t, NewBox(2+4, 2, 4+4, 4).Touching(NewBox(6, 5, 8, 7)))
	assert.True(t, NewBox(2+5, 2, 4+5, 4).Touching(NewBox(6, 5, 8, 7)))
	assert.True(t, NewBox(2+6, 2, 4+6, 4).Touching(NewBox(6, 5, 8, 7)))
	assert.True(t, NewBox(2+7, 2, 4+7, 4).Touching(NewBox(6, 5, 8, 7)))
	assert.False(t, NewBox(2+8, 2, 4+8, 4).Touching(NewBox(6, 5, 8, 7)))

	// b2 runs to the right of b1 with a gap of 1
	for y := 0; y < 9; y++ {
		assert.False(t, NewBox(6, 6-y, 8, 8-y).Touching(NewBox(2, 2, 4, 4)), y)
	}
	// b2 runs to the right of b1 with a gap of 0 (touching)
	assert.False(t, NewBox(5, 6-0, 8, 8-0).Touching(NewBox(2, 2, 4, 4)))
	assert.True(t, NewBox(5, 6-1, 8, 8-1).Touching(NewBox(2, 2, 4, 4)))
	assert.True(t, NewBox(5, 6-2, 8, 8-2).Touching(NewBox(2, 2, 4, 4)))
	assert.True(t, NewBox(5, 6-3, 8, 8-3).Touching(NewBox(2, 2, 4, 4)))
	assert.True(t, NewBox(5, 6-4, 8, 8-4).Touching(NewBox(2, 2, 4, 4)))
	assert.True(t, NewBox(5, 6-5, 8, 8-5).Touching(NewBox(2, 2, 4, 4)))
	assert.True(t, NewBox(5, 6-6, 8, 8-6).Touching(NewBox(2, 2, 4, 4)))
	assert.True(t, NewBox(5, 6-7, 8, 8-7).Touching(NewBox(2, 2, 4, 4)))
	assert.False(t, NewBox(5, 6-8, 8, 8-8).Touching(NewBox(2, 2, 4, 4)))
}
