package game

type Player struct {
	Health    int
	Hunger    int
	Thirst    int
	Inventory []string
}

func (p *Player) UpdateHealth(amount int) {
	p.Health += amount
	if p.Health < 0 {
		p.Health = 0
	}
}

func (p *Player) ManageHunger(amount int) {
	p.Hunger += amount
	if p.Hunger < 0 {
		p.Hunger = 0
	}
}

func (p *Player) ManageThirst(amount int) {
	p.Thirst += amount
	if p.Thirst < 0 {
		p.Thirst = 0
	}
}
