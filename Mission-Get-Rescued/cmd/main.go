package main

import (
    "fmt"
    "os"
    "Mission-Get-Rescued/internal/game"
)

func main() {
    fmt.Println("Welcome to Mission: Get Rescued!")

    engine := game.NewGameEngine()
    if err := engine.StartGame(); err != nil {
        fmt.Println("Error starting the game:", err)
        os.Exit(1)
    }

    for {
        if err := engine.UpdateGame(); err != nil {
            fmt.Println("Error updating the game:", err)
            break
        }
    }

    engine.EndGame()
    fmt.Println("Thank you for playing!")
}