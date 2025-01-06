package main

import "math"

type Path struct {
	MovementSpeed   float64
	FramesPerSecond int

	// DestX and DestY is where the player is moving towards. When a new
	// destination is set, we precalculate the Distance so that each FrameNumber
	// is proportionally correct to the distance of travel. Otherwise, the lower
	// resolution of screen would make the movement look jagged.
	DestX, DestY int
	IsMoving     bool
	Distance     float64
	FrameNumber  int

	// StartX and StartY is where the player was when the Dest was set. We need to
	// keep this value as PlayerX and PlayerY will change over multiple frames
	// during the journey and the correct new position of each frame.
	StartX, StartY int

	Boxer Boxer
}

func NewPath(startX, startY, destX, destY int, movementSpeed float64, fps int, boxer Boxer) *Path {
	return &Path{
		StartX:          startX,
		StartY:          startY,
		DestX:           destX,
		DestY:           destY,
		Distance:        distance(startX, startY, destX, destY),
		MovementSpeed:   movementSpeed,
		FramesPerSecond: fps,
		IsMoving:        true,
		Boxer:           boxer,
	}
}

func NewEmptyPath() *Path {
	return &Path{}
}

func (p *Path) Tick(canMove func(b Box) bool) (int, int) {
	newX, newY, isMoving := p.nextTick(p.FrameNumber + 1)
	if canMove(p.Boxer.Box(newX, newY)) {
		p.IsMoving = isMoving
		p.FrameNumber++
		return newX, newY
	}

	p.IsMoving = false
	newX, newY, _ = p.nextTick(p.FrameNumber)
	return newX, newY
}

func (p *Path) nextTick(frameNumber int) (int, int, bool) {
	// p.Distance is the total diagonal length to Dest. We calculate
	// `traveled` as the length of the diagonal based on movement speed
	// and how many frames have passed. From this ideal location we can
	// calculate the correct PlayerX and PlayerY.
	traveled := (p.MovementSpeed / float64(p.FramesPerSecond)) * float64(frameNumber)
	portion := traveled / p.Distance
	if portion > 1 {
		portion = 1
	}

	return p.StartX + int(math.Round(float64(p.DestX-p.StartX)*portion)),
		p.StartY + int(math.Round(float64(p.DestY-p.StartY)*portion)), portion < 1
}

func distance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x1-x2), 2) + math.Pow(float64(y1-y2), 2))
}

func distancef(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}
