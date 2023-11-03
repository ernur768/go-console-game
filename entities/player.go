package entities

import (
	gp "consoleTest/gamePhysics"
	"consoleTest/guns"
	"consoleTest/term"
)

const (
	playerSprite = "👾"
)

// Concrete player struct

type Player struct {
	id       int
	keyInput term.KeyInput
	position gp.Position
	gun      GunBehavior
}

type GunBehavior interface {
	Shoot()
}

// Constructor setter getter =======================================================

func NewPlayer(position gp.Position) *Player {
	player := &Player{
		keyInput: term.GetKeyInput(),
		position: position,
	}
	player.gun = guns.NewGun(player, 5) // 5 is the bullet speed
	return player
}

func (p *Player) SetGun(g GunBehavior) {
	p.gun = g
}

func (p *Player) SetID(ID int) {
	p.id = ID
}

func (p *Player) GetID() int {
	return p.id
}

func (p *Player) GetPosition() gp.Position {
	return p.position
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
