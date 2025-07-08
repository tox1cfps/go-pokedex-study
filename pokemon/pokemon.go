package pokemon

import "fmt"

type Pokemon struct {
	Name     string
	Type     string
	Attack   int
	Defense  int
	ID       int
	HP       int
	Captured bool
}

func (p Pokemon) GetInfo() string {
	return fmt.Sprintf("Name: %s\nType: %s\nID: %d\nHP: %d", p.Name, p.Type, p.ID, p.HP)
}

func (p *Pokemon) PokeAttack(p2 *Pokemon) string {
	if p.Attack > p2.Defense {
		p2.HP -= p.Attack
		if p2.HP <= 0 {
			p2.Captured = true
		}

		return p.Name + "Attacked" + p2.Name

	} else if p.Attack < p2.Defense {
		return p2.Name + "Attacked" + p.Name
	} else {
		return "It's a Tie between" + p.Name + "and" + p2.Name
	}
}

func (p Pokemon) Capture(p2 *Pokemon) string {
	if p2.HP == 0 {
		p2.Captured = true
		return p2.Name + "Was captured by" + p.Name + "!"
	} else {
		return "Capture failed\nThe battle continues... "
	}
}
