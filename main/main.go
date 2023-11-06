package main

import (
	"consoleTest/entities"
	gp "consoleTest/gamePhysics"
	"consoleTest/term"
	"time"
)

func main() {

	gameStart()
	gameLoop()

}

// Before game starts ===============================

func gameStart() {
	term.HideCursor()

	Entities := gp.GetEntities()
	//Asteroids := entities.GetCopyAsteroids()

	health := entities.NewPlayerHealth()
	gp.AppendEntity(health)

	spacecraft := entities.NewSpacecraft()
	health.RegisterObserver(spacecraft)
	gp.AppendEntity(spacecraft)

	spawner := entities.NewAsteroidSpawner(health)
	health.RegisterObserver(spawner)
	gp.AppendEntity(spawner)

	player := entities.NewPlayer(gp.NewPosition(30, 14))
	gp.AppendEntity(player)

	for _, object := range *Entities {
		object.Start()
	}

}

// Game loop =======================================

func gameLoop() {
	Entities := gp.GetCopy()
	Asteroids := entities.GetCopyAsteroids()

	for {
		term.Clear()

		for _, entity := range Entities {
			entity.Update()
		}
		for _, asteroid := range Asteroids {
			asteroid.Update()
		}
		Entities = gp.GetCopy()
		Asteroids = entities.GetCopyAsteroids()

		term.Render()
		time.Sleep(50 * time.Millisecond)
	}
}

// singleton +
// strategy +
// observer
// factory method
// abstract factory
// decorator
