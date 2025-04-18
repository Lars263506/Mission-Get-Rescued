package utils

import (
	"math/rand"
	"strings"
	"time"
)

// GenerateRandomEvent returns a random event based on the provided events slice.
func GenerateRandomEvent(events []string) string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return events[rng.Intn(len(events))]
}

// LogMessage formats and logs a message to the console.
func LogMessage(message string) {
	formattedMessage := strings.ToUpper(message)
	println(formattedMessage)
}

// ParseInput processes user input and returns a trimmed string.
func ParseInput(input string) string {
	return strings.TrimSpace(input)
}
