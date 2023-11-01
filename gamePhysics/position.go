package gamePhysics

const (
	boundX = 160
	boundY = 28
)

type Position struct {
	X int
	Y int
}

func NewPosition(x, y int) Position {
	return Position{x, y}
}

func (p *Position) StepUp() {
	if p.Y <= 1 {
		return
	}
	p.Y = p.Y - 1
}

func (p *Position) StepDown() {
	if p.Y >= boundY {
		return
	}
	p.Y = p.Y + 1
}

func (p *Position) StepLeft() {
	if p.X <= 1 {
		return
	}
	p.X = p.X - 1
}

func (p *Position) StepRight() {
	if p.X >= boundX {
		return
	}
	p.X = p.X + 1
}
