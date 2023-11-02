package entities

import (
	gp "consoleTest/gamePhysics"
	"consoleTest/term"
)

var sprite = []string{
	"              ======\\",
	"             /       \\",
	"------------/        /",
	"===========/         |",
	"                     |",
	"             ||      |",
	"        ======\\\\\\    |",
	"===============\\\\\\\\   \\",
	"===============////   /",
	"       =======///    |",
	"             ||      |",
	"                     |",
	"===========\\         |",
	"------------\\        \\",
	"             \\       /",
	"              ======/",
}

type Spacecraft struct {
	id       int
	position gp.Position
}

// getters and setters ========================

func NewSpacecraft() *Spacecraft {
	return &Spacecraft{
		position: gp.NewPosition(0, 7),
	}
}

func (s *Spacecraft) SetID(ID int) {
	s.id = ID
}

func (s *Spacecraft) GetID() int {
	return s.id
}

func (s *Spacecraft) GetPosition() gp.Position {
	return s.position
}

// Entity methods ==============================================

func (s *Spacecraft) Start() {
	//TODO implement me
}

func (s *Spacecraft) Update() {
	s.drowSprite()
}

func (s *Spacecraft) Finalize() {
	//TODO implement me
	panic("implement me")
}

func (s *Spacecraft) drowSprite() {
	for _, row := range sprite {
		term.MoveCursorAndDraw(s.position, row)
		s.position.StepDown()
	}
	s.position.Y = 7
}
