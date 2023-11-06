package entities

import (
	gp "consoleTest/gamePhysics"
	"time"
)

type GunBehavior interface {
	Shoot()
	RenderGunName()
}

func NewGun(owner *Player, bulletSpeed int, gunType GunType) GunBehavior {
	switch gunType {
	case PistolType:
		return &Pistol{owner: owner, bulletSpeed: bulletSpeed, lastShot: time.Now()}
	case RifleType:
		return &Rifle{owner: owner, bulletSpeed: bulletSpeed, lastShot: time.Now()}
	case MiniGunType:
		return &MiniGun{owner: owner, bulletSpeed: bulletSpeed, lastShot: time.Now()}
	default:
		return nil
	}
}

type GunType int

const (
	PistolType GunType = iota
	RifleType
	MiniGunType
)

type Pistol struct {
	owner       *Player
	bulletSpeed int
	lastShot    time.Time
	pos         gp.Position
}

func (p *Pistol) Shoot() {
	if time.Since(p.lastShot) < time.Minute/30 {
		return
	}
	p.lastShot = time.Now()
	bullet := NewBullet(p.owner.GetPosition(), p.bulletSpeed)
	bullet.Start()
	gp.AppendEntity(bullet)
}

func (p *Pistol) Render() {
	return
}

type Rifle struct {
	owner       *Player
	bulletSpeed int
	lastShot    time.Time
}

func (g *Rifle) Shoot() {
	if time.Since(g.lastShot) < time.Minute/100 {
		return
	}
	g.lastShot = time.Now()
	pos1 := g.owner.position
	pos1.Y = pos1.Y + 1
	pos3 := g.owner.position
	pos3.Y = pos3.Y - 1
	upBullet := NewBullet(pos1, g.bulletSpeed)
	centerBullet := NewBullet(g.owner.position, g.bulletSpeed)
	downBullet := NewBullet(pos3, g.bulletSpeed)
	upBullet.Start()
	centerBullet.Start()
	downBullet.Start()
	gp.AppendEntity(upBullet)
	gp.AppendEntity(centerBullet)
	gp.AppendEntity(downBullet)
}

type MiniGun struct {
	owner       *Player
	bulletSpeed int
	lastShot    time.Time
}

func (m *MiniGun) Shoot() {
	if time.Since(m.lastShot) < time.Second/10 {
		return
	}
	m.lastShot = time.Now()
	bullet := NewBullet(m.owner.GetPosition(), m.bulletSpeed)
	bullet.Start()
	gp.AppendEntity(bullet)
}
