package entities

import (
	gp "consoleTest/gamePhysics"
	"math/rand"
	"time"
)

type AsteroidSpawner struct {
	id            int
	position      gp.Position
	pauseDuration int
	asteroidSpeed int
	destroyed     chan struct{}
	playerHealth  *PlayerHealth
}

func NewAsteroidSpawner(health *PlayerHealth) *AsteroidSpawner {
	return &AsteroidSpawner{
		position:      gp.NewPosition(0, 0),
		pauseDuration: 2,
		asteroidSpeed: 50,
		destroyed:     make(chan struct{}, 1),
		playerHealth:  health,
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
	//
}

func (as *AsteroidSpawner) Finalize() {
	as.destroyed <- struct{}{}
}

func (as *AsteroidSpawner) spawnAsteroid() {
	pos := gp.NewPosition(159, rand.Int()%28+1)
	asteroid := NewAsteroid(pos, as.asteroidSpeed, as.playerHealth)
	AppendAsteroid(asteroid)
	asteroid.Start()
	time.Sleep(time.Duration(as.pauseDuration) * time.Second)
}
