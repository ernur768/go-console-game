package entities

import (
	gp "consoleTest/gamePhysics"
)

type Gun struct {
	owner       *Player
	bulletSpeed int
}

func NewGun(owner *Player, bulletSpeed int) *Gun {
	return &Gun{
		owner:       owner,
		bulletSpeed: bulletSpeed,
	}
}

func (g *Gun) Shoot() {
	bullet := NewBullet(g.owner.GetPosition(), g.bulletSpeed)
	bullet.Start()
	gp.AppendEntity(bullet)
}
