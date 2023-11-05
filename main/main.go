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

	spacecraft := entities.NewSpacecraft()
	gp.AppendEntity(spacecraft)

	health := entities.NewPlayerHealth()
	health.RegisterObserver(spacecraft)
	gp.AppendEntity(health)

	player := entities.NewPlayer(gp.NewPosition(5, 5))
	gp.AppendEntity(player)

	spawner := entities.NewAsteroidSpawner(health)
	gp.AppendEntity(spawner)

	//a1 := entities.NewAsteroid(gp.NewPosition(60, 10), 50, health)
	//entities.AppendAsteroid(a1)
	//
	//a2 := entities.NewAsteroid(gp.NewPosition(70, 20), 50, health)
	//entities.AppendAsteroid(a2)
	//
	//a3 := entities.NewAsteroid(gp.NewPosition(75, 15), 50, health)
	//entities.AppendAsteroid(a3)
	//
	//a4 := entities.NewAsteroid(gp.NewPosition(70, 17), 50, health)
	//entities.AppendAsteroid(a4)

	for _, object := range *Entities {
		object.Start()
	}

	//Asteroids = entities.GetCopyAsteroids()
	//for _, asteroid := range Asteroids {
	//	asteroid.Start()
	//}

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
