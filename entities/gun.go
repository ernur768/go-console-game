package entities

import (
	gp "consoleTest/gamePhysics"
	"consoleTest/term"
	"time"
)

type GunBehavior interface {
	Shoot()
	RenderGunName()
}

func NewGun(owner *Player, bulletSpeed int, gunType GunType) GunBehavior {
	switch gunType {
	case PistolType:
		return &Pistol{owner: owner, bulletSpeed: bulletSpeed, lastShot: time.Now(), pos: gp.NewPosition(50, 1)}
	case RifleType:
		return &Rifle{owner: owner, bulletSpeed: bulletSpeed, lastShot: time.Now(), pos: gp.NewPosition(50, 1)}
	case MiniGunType:
		return &MiniGun{owner: owner, bulletSpeed: bulletSpeed, lastShot: time.Now(), pos: gp.NewPosition(50, 1)}
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
	if time.Since(p.lastShot) < time.Minute/70 {
		return
	}
	p.lastShot = time.Now()
	bullet := NewBullet(p.owner.GetPosition(), p.bulletSpeed)
	bullet.Start()
	gp.AppendEntity(bullet)
}

func (p *Pistol) RenderGunName() {
	term.MoveCursorAndDraw(p.pos, "Pistol")
}

type Rifle struct {
	owner       *Player
	bulletSpeed int
	lastShot    time.Time
	pos         gp.Position
}

func (r *Rifle) Shoot() {
	if time.Since(r.lastShot) < time.Minute/40 {
		return
	}
	r.lastShot = time.Now()
	pos1 := r.owner.position
	pos1.Y = pos1.Y + 1
	pos3 := r.owner.position
	pos3.Y = pos3.Y - 1
	upBullet := NewBullet(pos1, r.bulletSpeed)
	centerBullet := NewBullet(r.owner.position, r.bulletSpeed)
	downBullet := NewBullet(pos3, r.bulletSpeed)
	upBullet.Start()
	centerBullet.Start()
	downBullet.Start()
	gp.AppendEntity(upBullet)
	gp.AppendEntity(centerBullet)
	gp.AppendEntity(downBullet)
}

func (r *Rifle) RenderGunName() {
	term.MoveCursorAndDraw(r.pos, "Rifle")
}

type MiniGun struct {
	owner       *Player
	bulletSpeed int
	lastShot    time.Time
	pos         gp.Position
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

func (m *MiniGun) RenderGunName() {
	term.MoveCursorAndDraw(m.pos, "MiniGun")
}
