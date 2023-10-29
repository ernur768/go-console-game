package gamePhysics

type Position struct {
	X int
	Y int
}

func NewPosition(x, y int) Position {
	return Position{x, y}
}

func (p *Position) StepUp() {
	p.Y = p.Y - 1
}

func (p *Position) StepDown() {
	p.Y = p.Y + 1
}

func (p *Position) StepLeft() {
	p.X = p.X - 1
}

func (p *Position) StepRight() {
	p.X = p.X + 1
}
