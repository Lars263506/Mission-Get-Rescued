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

func (g *GameEngine) StartGame() {
    g.isRunning = true
    fmt.Println("Game started. Survive until rescued!")
}

func (g *GameEngine) UpdateGame() {
    if g.isRunning {
        // Update game state logic here
        fmt.Println("Updating game state...")
    }
}

func (g *GameEngine) EndGame() {
    g.isRunning = false
    fmt.Println("Game over. You did not survive.")
}