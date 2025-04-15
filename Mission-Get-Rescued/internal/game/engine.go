package game

import "fmt"

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
}
