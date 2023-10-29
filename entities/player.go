package entities

import (
	gp "consoleTest/gamePhysics"
	"consoleTest/term"
)

const (
	playerSprite = "👾"
)

// Player's gun interface

type PlayerGun interface {
	Shoot()
}

// Concrete player struct

type Player struct {
	id       int
	keyInput term.KeyInput
	position gp.Position
	gun      PlayerGun
}

// Constructor setter getter =======================================================

func NewPlayer(position gp.Position) *Player {
	return &Player{
		keyInput: term.GetKeyInput(),
		position: position,
	}
}

func (p *Player) SetGun(g PlayerGun) {
	p.gun = g
}

func (p *Player) SetID(ID int) {
	p.id = ID
}

func (p *Player) GetID() int {
	return p.id
}

// Entity interface implementations ===================================

func (p *Player) Start() {
	go p.movement()
}

func (p *Player) Update() {
	term.MoveCursorAndDraw(p.position, playerSprite)
}

func (p *Player) Finalize() {
	p.keyInput.Close()
}

func (p *Player) movement() {
	for key := range p.keyInput.C {
		switch key {
		case 'w':
			p.position.StepUp()
		case 'a':
			p.position.StepLeft()
		case 's':
			p.position.StepDown()
		case 'd':
			p.position.StepRight()
		case 'j':
			p.gun.Shoot()
		case 'i':
			gp.DestroyEntity(p)
		}
	}
}
