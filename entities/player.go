package entities

import (
	gp "consoleTest/gamePhysics"
	"consoleTest/term"
)

const (
	playerSprite = "👾"
)

type Player struct {
	id       int
	keyInput term.KeyInput
	position gp.Position
	gun      GunBehavior
}

func NewPlayer(position gp.Position) *Player {
	player := &Player{
		keyInput: term.GetKeyInput(),
		position: position,
	}
	player.gun = NewGun(player, 10, PistolType)
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

func (p *Player) Start() {
	go p.movement()
}

func (p *Player) Update() {
	term.MoveCursorAndDraw(p.position, playerSprite)
}

func (p *Player) Finalize() {
	p.keyInput.Close()
}

func (p *Player) changeToPistol() {
	p.gun = NewGun(p, 10, PistolType)
}

func (p *Player) changeToRifle() {
	p.gun = NewGun(p, 10, RifleType)
}

func (p *Player) changeToMiniGun() {
	p.gun = NewGun(p, 10, MiniGunType)
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
		case '1':
			p.changeToPistol()
		case '2':
			p.changeToRifle()
		case '3':
			p.changeToMiniGun()
		}
	}
}
