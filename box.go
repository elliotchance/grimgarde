package main

type Boxer interface {
	Box(offsetX, offsetY int) Box
}

type Box struct {
	x1, y1, x2, y2 int
}

// NewBox creates a new Box from two points (bottom-left and top-right)
func NewBox(x1, y1, x2, y2 int) Box {
	return Box{x1, y1, x2, y2}
}

func (b1 Box) Intersect(b2 Box) bool {
	// +1/-1 here is to adjust for the box itself having a width of 1.
	// Calculate the maximum of the bottom-left x and y coordinates
	ix1 := max(b1.x1-2, b2.x1-1)
	iy1 := max(b1.y1-2, b2.y1-1)
	// Calculate the minimum of the top-right x and y coordinates
	ix2 := min(b1.x2+1, b2.x2)
	iy2 := min(b1.y2+1, b2.y2)

	// If there's no intersection (if one box is to the left, right, above, or
	// below the other), return false
	return ix1 < ix2 && iy1 < iy2
}

// Distance is the length between the two centres of the boxes.
func (b1 Box) Distance(b2 Box) float64 {
	return distancef(middle(b1.x1, b1.x2), middle(b1.y1, b1.y2),
		middle(b2.x1, b2.x2), middle(b2.y1, b2.y2))
}

func (b1 Box) Touching(b2 Box) bool {
	touchingLeft := b1.x2 == b2.x1-1 && ((b1.y1+2 >= b2.y1-1) && (b1.y1 <= b2.y2+1))
	touchingRight := b1.x1-1 == b2.x2 && ((b1.y1+2 >= b2.y1-1) && (b1.y1 <= b2.y2+1))
	touchingTop := b1.y2 == b2.y1-1 && ((b1.x1+2 >= b2.x1-1) && (b1.x1 <= b2.x2+1))
	touchingBottom := b1.y1-1 == b2.y2 && ((b1.x1+2 >= b2.x1-1) && (b1.x1 <= b2.x2+1))

	return touchingLeft || touchingRight || touchingTop || touchingBottom
}

func (b1 Box) Contains(x, y int) bool {
	return x >= b1.x1 && x <= b1.x2 && y >= b1.y1 && y <= b1.y2
}

func middle(a, b int) float64 {
	return (float64(a) + float64(b)) / 2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
