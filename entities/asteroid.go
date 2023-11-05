package entities

import (
	gp "consoleTest/gamePhysics"
	"consoleTest/term"
	"time"
)

type Asteroid struct {
	id           int
	sprite       string
	position     gp.Position
	pauseSpeed   int
	playerHealth *PlayerHealth
	destroyed    chan struct{}
}

func NewAsteroid(pos gp.Position, pauseSpeed int, playerHealth *PlayerHealth) *Asteroid {
	return &Asteroid{
		sprite:       "◌",
		position:     pos,
		playerHealth: playerHealth,
		pauseSpeed:   pauseSpeed,
		destroyed:    make(chan struct{}, 1),
	}
}

// getters and setters =====================

func (a *Asteroid) SetID(ID int) {
	a.id = ID
}

func (a *Asteroid) GetID() int {
	return a.id
}

func (a *Asteroid) GetPosition() gp.Position {
	return a.position
}

// Entity methods =====================================

func (a *Asteroid) Start() {
	//fmt.Println("started")
	go a.movement()
}

func (a *Asteroid) Update() {
	if a.position.X == 1 {
		DestroyAsteroid(*a)
		return
	}
	if a.position.X <= 23 && a.position.Y >= 7 && a.position.Y <= 22 {
		a.playerHealth.GetHit()
		DestroyAsteroid(*a)
		return
	}
	term.MoveCursorAndDraw(a.position, a.sprite)
}

func (a *Asteroid) Finalize() {
	a.destroyed <- struct{}{}
}

func (a *Asteroid) movement() {
	for {
		select {
		case <-a.destroyed:
			return
		default:
			a.position.StepLeft()
			time.Sleep(time.Duration(a.pauseSpeed) * time.Millisecond)
		}
	}
}
