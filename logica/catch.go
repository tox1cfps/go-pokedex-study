package logica

func (p Pokemon) Capture(p2 *Pokemon) string {
	if p2.HP == 0 {
		p2.Captured = true
		return p2.Name + "Was captured by" + p.Name + "!"
	} else {
		return "The battle continues... "
	}
}
