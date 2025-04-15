package game

func GetAllEvents() []Event {
	return []Event{
		{
			Description: "You find spoiled food!",
			Impact: func(p *Player) {
				p.UpdateHealth(-10)
				p.ManageHunger(-5)
			},
			Condition: func(p *Player) bool {
				return p.Hunger < 30 // Trigger if hunger is low
			},
			Priority: 2,
		},
		{
			Description: "You drink contaminated water!",
			Impact: func(p *Player) {
				p.UpdateHealth(-15)
				p.ManageThirst(-5)
			},
			Condition: func(p *Player) bool {
				return p.Thirst < 20 // Trigger if thirst is low
			},
			Priority: 3,
		},
		{
			Description: "A guard provides you with a rotten apple, it looks suspicious!",
			Impact: func(p *Player) {
				p.UpdateHealth(-5)
				p.ManageHunger(+12)
				p.ManageThirst(+8)
			},
			Condition: func(p *Player) bool {
				return p.Hunger < 10
			},
			Priority: 2,
		},
		{
			Description: "You cannot find a way out, your hope is fading!",
			Impact: func(p *Player) {
				p.UpdateHealth(-5)
				p.ManageHunger(-5)
				p.ManageThirst(-5)
			},
			Condition: func(p *Player) bool {
				return true // Always eligible to trigger
			},
			Priority: 1,
		},
	}
}
