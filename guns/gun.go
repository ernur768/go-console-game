package guns

import (
	"consoleTest/entities"
	gp "consoleTest/gamePhysics"
)

type Gun struct {
	owner       *entities.Player
	bulletSpeed int
}

func NewGun(owner *entities.Player, bulletSpeed int) *Gun {
	return &Gun{
		owner:       owner,
		bulletSpeed: bulletSpeed,
	}
}

func (g *Gun) Shoot() {
	bullet := entities.NewBullet(g.owner.GetPosition(), g.bulletSpeed)
	gp.AppendEntity(bullet)
}
