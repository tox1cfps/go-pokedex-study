package pokemon

import "fmt"

type Pokemon struct {
	Name string
	Type string
	ID   int
}

func (p Pokemon) GetInfo() string {
	return fmt.Sprintf("Name: %s\nType: %s\nID: %d", p.Name, p.Type, p.ID)
}
