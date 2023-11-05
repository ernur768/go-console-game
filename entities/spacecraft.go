package entities

import (
	gp "consoleTest/gamePhysics"
	"consoleTest/term"
)

var s = []string{
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
	"                    |",
	"===========\\         |",
	"------------\\        \\",
	"             \\       /",
	"              ======/",
}

// ¤ ⊗ ⊝ ⨳ ⨠ ⩩ ⩨ ⩸ ϡ Ϡ ⁜ ×

type Spacecraft struct {
	id       int
	position gp.Position
	sprite   []string
	health   int
}

// getters and setters ========================

func NewSpacecraft() *Spacecraft {
	return &Spacecraft{
		position: gp.NewPosition(0, 7),
		sprite:   s,
		health:   5,
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

// Spacecraft methods ====================================

func (s *Spacecraft) drowSprite() {
	for _, row := range s.sprite {
		term.MoveCursorAndDraw(s.position, row)
		s.position.StepDown()
	}
	s.position.Y = 7
}

// Observer method ================================

func (s *Spacecraft) PlayerGotHit() {
	s.health--
	if s.health == 4 {
		s.sprite[2] = "------------/    ⩩   /"
		s.sprite[3] = "===========/     ×⩨  |"
	} else if s.health == 3 {
		s.sprite[11] = "               ¤Ϡ    |"
		s.sprite[12] = "===========\\  ϡ      |"
	} else if s.health == 2 {
		s.sprite[8] = "===========▩===////   /"
		s.sprite[13] = "    Ϡ-------\\        \\"
	} else if s.health == 1 {
		s.sprite[1] = "              \\      \\"
		s.sprite[2] = "              /      /"
		s.sprite[5] = "      -⨀-    ||      |"
	}
}
