package game

import "fmt"

// Event represents an in-game event that can affect the player's survival.
type Event struct {
	Description string
	Impact      func(*Player)
}

// HandleEvent processes an event and applies its impact on the player.
func HandleEvent(event Event, player *Player) {
	fmt.Println(event.Description)
	event.Impact(player)
}

// CheckSurvival evaluates the player's current status to determine if they are still alive.
func CheckSurvival(player *Player) bool {
	return player.Health > 0 && player.Hunger < 100 && player.Thirst < 100
}

// ExecuteAction performs a specific action based on player input and updates the game state.
func ExecuteAction(action string, player *Player) {
	switch action {
	case "eat":
		player.ManageHunger(-10)
		fmt.Println("You eat some food.")
	case "drink":
		player.ManageThirst(-10)
		fmt.Println("You drink some water.")
	default:
		fmt.Println("Action not recognized.")
	}
}