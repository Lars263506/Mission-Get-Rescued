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
	figure := figure.NewFigure("Mission: Get Rescued", "", true)
	figure.Print()

	fmt.Println("Welcome to Mission: Get Rescued!")

	engine := game.NewGameEngine()
	engine.StartGame()

	// Define the player outside the loop
	player := game.Player{
		Health: 100,
		Hunger: 50,
		Thirst: 50,
	}

	for engine.IsRunningGame() {
		var action string
		fmt.Println("Choose an action (eat/drink):")
		fmt.Scanln(&action)
		clearConsole()

		// Update the player's state based on the action
		game.ExecuteAction(action, &player)

		// Trigger an event based on conditions
		game.TriggerEvent(&player)

		// Check if the player is still alive
		if !game.CheckSurvival(&player) {
			engine.EndGame()
			fmt.Println("You have died. Game over.")
			break
		}

		// Print the player's current status
		fmt.Printf("Player Status - Health: %d, Hunger: %d, Thirst: %d\n", player.Health, player.Hunger, player.Thirst)
		fmt.Println("Survive until rescued!")
	}

	engine.EndGame()
	fmt.Println("Thank you for playing!")
}
