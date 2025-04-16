package main

import (
	"Mission-Get-Rescued/internal/game"
	"fmt"
	"os"
	"os/exec"

	"github.com/common-nighthawk/go-figure"
)

func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls") // Windows-specific command
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	titleFigure := figure.NewFigure("Mission: Get Rescued", "", true)
	titleFigure.Print()

	engine := game.NewGameEngine()
	engine.StartGame()

	// Define the player outside the loop
	player := game.Player{
		Health: 100,
		Hunger: 50,
		Thirst: 50,
	}

	for engine.IsRunningGame() {
		clearConsole() // Clear the console at the start of each iteration

		figure := figure.NewFigure("Mission: Get Rescued", "", true)
		figure.Print()

		fmt.Println("-----------------------------------------------------------------------------------------------------------------------------")

		// Print the player's current status
		fmt.Printf("Player Status - Health: %d, Hunger: %d, Thirst: %d\n", player.Health, player.Hunger, player.Thirst)

		// Trigger an event based on conditions
		game.TriggerEvent(&player)

		// Check if the player is still alive
		if !game.CheckSurvival(&player) {
			engine.EndGame()
			fmt.Println("You have died. Game over.")
			break
		}

		// Prompt the player for the next action
		fmt.Println("Choose an action for the next turn (eat/drink):")
		var action string
		fmt.Scanln(&action)

		// Update the player's state based on the action
		game.ExecuteAction(action, &player)
	}

	engine.EndGame()
	fmt.Println("Thank you for playing!")
}
