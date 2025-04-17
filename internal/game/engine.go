package game

import (
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

type GameEngine struct {
	playerStatus string
	events       []string
	isRunning    bool
}

func NewGameEngine() *GameEngine {
	return &GameEngine{
		playerStatus: "alive",
		events:       []string{},
		isRunning:    false,
	}
}

func (e *GameEngine) StartGame() {
	e.isRunning = true
	fmt.Println("Game started. Survive until rescued!")
}

func (e *GameEngine) IsRunningGame() bool {
	return e.isRunning
}

func (e *GameEngine) EndGame() {
	e.isRunning = false
	fmt.Println("Game over. You did not survive.")

	for {
		fmt.Println("Choose an action: restart or exit")
		var action string
		fmt.Scanln(&action)

		switch action {
		case "restart":
			fmt.Println("Restarting the game...")
			RunGame()
			return
		case "exit":
			fmt.Println("Exiting the game. Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid action. Please choose 'restart' or 'exit'.")
		}
	}
}

func RunGame() {
	engine := NewGameEngine()
	engine.StartGame()

	// Define the player outside the loop
	player := Player{
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
		TriggerEvent(&player)

		// Check if the player is still alive
		if !CheckSurvival(&player) {
			engine.EndGame()
			fmt.Println("You have died. Game over.")
			return
		}

		// Prompt the player for the next action
		fmt.Println("Choose an action for the next turn (eat/drink):")
		var action string
		fmt.Scanln(&action)

		// Update the player's state based on the action
		ExecuteAction(action, &player)
	}

	engine.EndGame()
	fmt.Println("Thank you for playing!")
}
