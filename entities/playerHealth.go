package entities

import (
	gp "consoleTest/gamePhysics"
	"consoleTest/term"
	"strings"
)

type PlayerHealth struct {
	sprite       string
	id           int
	position     gp.Position
	healthPoints int
}

func NewPlayerHealth() *PlayerHealth {
	return &PlayerHealth{
		sprite:       "♥ ",
		position:     gp.NewPosition(75, 1),
		healthPoints: 5,
	}
}

// getters and setters =====================

func (ph *PlayerHealth) SetID(ID int) {
	ph.id = ID
}

func (ph *PlayerHealth) GetID() int {
	return ph.id
}

func (ph *PlayerHealth) GetPosition() gp.Position {
	return ph.position
}

func (ph *PlayerHealth) Start() {
	//TODO implement me
}

func (ph *PlayerHealth) Update() {
	term.MoveCursorAndDraw(ph.position, strings.Repeat(ph.sprite, ph.healthPoints))
}

func (ph *PlayerHealth) Finalize() {
	//TODO implement me
	panic("implement me")
}

func (ph *PlayerHealth) GetHit() {
	ph.healthPoints--
}
