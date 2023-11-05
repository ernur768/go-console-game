package entities

import (
	gp "consoleTest/gamePhysics"
	"consoleTest/term"
	"strings"
)

// used interface ===========================================

type Observer interface {
	PlayerGotHit()
}

// PlayerHealth struct =======================================

type PlayerHealth struct {
	sprite       string
	id           int
	position     gp.Position
	healthPoints int
	observers    []Observer
}

func NewPlayerHealth() *PlayerHealth {
	return &PlayerHealth{
		sprite:       "♥ ",
		position:     gp.NewPosition(75, 1),
		healthPoints: 5,
		observers:    make([]Observer, 0),
	}
}

// getters and setters =============================================

func (ph *PlayerHealth) SetID(ID int) {
	ph.id = ID
}

func (ph *PlayerHealth) GetID() int {
	return ph.id
}

func (ph *PlayerHealth) GetPosition() gp.Position {
	return ph.position
}

// Entity methods ======================================================

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

// playerHealth methods ======================================================

func (ph *PlayerHealth) GetHit() {
	ph.healthPoints--
	ph.notifyObservers()
}

// Observer methods =========================================================

func (ph *PlayerHealth) RegisterObserver(o Observer) {
	ph.observers = append(ph.observers, o)
}

func (ph *PlayerHealth) notifyObservers() {
	for _, observer := range ph.observers {
		observer.PlayerGotHit()
	}
}
