package entities

import (
	gp "consoleTest/gamePhysics"
	"math/rand"
	"time"
)

type AsteroidSpawner struct {
	id                 int
	position           gp.Position
	pauseDuration      int
	asteroidPauseSpeed int
	destroyed          chan struct{}
	playerHealth       *PlayerHealth
}

func NewAsteroidSpawner(health *PlayerHealth) *AsteroidSpawner {
	return &AsteroidSpawner{
		position:           gp.NewPosition(0, 0),
		pauseDuration:      5,
		asteroidPauseSpeed: 50,
		destroyed:          make(chan struct{}, 1),
		playerHealth:       health,
	}
}

// getters and setters =====================

func (as *AsteroidSpawner) SetID(ID int) {
	as.id = ID
}

func (as *AsteroidSpawner) GetID() int {
	return as.id
}

func (as *AsteroidSpawner) GetPosition() gp.Position {
	return as.position
}

// Entity methods =====================================

func (as *AsteroidSpawner) Start() {
	go func() {
		for {
			select {
			case <-as.destroyed:
				return
			default:
				as.spawnAsteroid()
			}
		}
	}()
}

func (as *AsteroidSpawner) Update() {
	if as.playerHealth.healthPoints <= 1 {
		gp.DestroyEntity(as)
	}
}

func (as *AsteroidSpawner) Finalize() {
	as.destroyed <- struct{}{}
}

func (as *AsteroidSpawner) spawnAsteroid() {
	pos := gp.NewPosition(159, rand.Int()%21+4)
	asteroid := NewAsteroid(pos, as.asteroidPauseSpeed, as.playerHealth)
	AppendAsteroid(asteroid)
	asteroid.Start()
	time.Sleep(time.Duration(as.pauseDuration) * time.Second)
}

func (as *AsteroidSpawner) PlayerGotHit() {
	if as.pauseDuration <= 2 {
		return
	}
	as.pauseDuration -= 1
}
