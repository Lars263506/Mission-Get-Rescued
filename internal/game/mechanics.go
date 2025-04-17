package game

import (
	"fmt"
)

// Event represents an in-game event that can affect the player's survival.
type Event struct {
	Description string
	Impact      func(*Player)
	Condition   func(*Player) bool
	Priority    int
}

// HandleEvent processes an event and applies its impact on the player.
func HandleEvent(event Event, player *Player) {
	fmt.Println(event.Description)
	event.Impact(player)
}

// TriggerEvent selects and triggers the highest priority event that meets its condition.
func TriggerEvent(player *Player) {
	events := GetAllEvents()
	var selectedEvent *Event
	highestPriority := -1

	for _, event := range events {
		if event.Condition(player) && event.Priority > highestPriority {
			selectedEvent = &event
			highestPriority = event.Priority
		}
	}

	if selectedEvent != nil {
		HandleEvent(*selectedEvent, player)
	}
}

// CheckSurvival evaluates the player's current status to determine if they are still alive.
func CheckSurvival(player *Player) bool {
	return player.Health > 0 && player.Hunger > 0 && player.Thirst > 0
}

// ExecuteAction performs a specific action based on player input and updates the game state.
func ExecuteAction(action string, player *Player) {
	switch action {
	case "eat":
		player.ManageHunger(+10)
		fmt.Println("You eat some food.")
	case "drink":
		player.ManageThirst(+10)
		fmt.Println("You drink some water.")
	default:
		fmt.Println("Action not recognized.")
	}
}
