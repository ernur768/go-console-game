package entities

import (
	gp "consoleTest/gamePhysics"
	"consoleTest/term"
	"time"
)

type Bullet struct {
	id        int
	sprite    string
	position  gp.Position
	speed     int
	destroyed chan struct{}
}

func NewBullet(startPosition gp.Position, speed int) *Bullet {
	return &Bullet{
		sprite:    "->",
		position:  startPosition,
		speed:     speed,
		destroyed: make(chan struct{}, 1),
	}
}

func (b *Bullet) SetID(ID int) {
	b.id = ID
}

func (b *Bullet) GetID() int {
	return b.id
}

func (b *Bullet) GetPosition() gp.Position {
	return b.position
}

func (b *Bullet) Start() {
	go b.movement()
}

func (b *Bullet) movement() {
	for {
		select {
		case <-b.destroyed:
			return
		default:
			b.position.StepRight()
			time.Sleep(time.Duration(b.speed) * time.Millisecond)
		}
	}
}

func (b *Bullet) Update() {
	if b.position.X >= 155 {
		gp.DestroyEntity(b)
		return
	}
	b.Collision()
	term.MoveCursorAndDraw(b.position, b.sprite)
}

func (b *Bullet) Finalize() {
	b.destroyed <- struct{}{}
}

func (b *Bullet) Collision() {

	Asteroids := GetCopyAsteroids()
	for _, a := range Asteroids {
		if a.position.Y == b.position.Y && a.position.X+1 <= b.position.X {
			gp.DestroyEntity(b)
			DestroyAsteroid(*a)
			return
		}
	}
}
