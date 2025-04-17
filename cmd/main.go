package main

import (
	"Mission-Get-Rescued/internal/game"
	"os/exec"
)

func setFullscreen() {
	cmd := exec.Command("powershell", "-Command", "Add-Type -AssemblyName System.Windows.Forms; [System.Windows.Forms.SendKeys]::SendWait('%{ENTER}')")
	cmd.Run()
}

func main() {
	setFullscreen() // Set the console to fullscreen mode
	game.RunGame()  // Call the reusable game logic
}
