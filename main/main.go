package main

import (
	"consoleTest/entities"
	gp "consoleTest/gamePhysics"
	"consoleTest/term"
)

func main() {

	gameStart()
	gameLoop()

}

// Before game starts ===============================

func gameStart() {
	term.HideCursor()

	Entities := gp.GetEntities()

	player := entities.NewPlayer(gp.NewPosition(5, 5))
	gp.AppendEntity(player)

	//player2 := entities.NewPlayer(gp.NewPosition(10, 10))
	//gp.AppendEntity(player2)

	for _, object := range *Entities {
		object.Start()
	}
}

// Game loop =======================================

func gameLoop() {
	Entities := gp.GetEntities()

	for {
		term.Clear()
		for _, entity := range *Entities {
			entity.Update()
		}

		term.Render()
	}
}

// singleton +
// strategy +
// observer
// factory method
// abstract factory
// decorator