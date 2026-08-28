package editorial

type Gate struct {
	Name     string
	Required bool
	Check    func(string) bool
}

func DefaultGates() []Gate {
	return []Gate{
		{"length", true, func(s string) bool { return len(s) >= 120 }},
		{"paragraphs", false, func(s string) bool { return len(s) > 0 }},
		{"disclosure", true, func(s string) bool { return len(s) > 20 }},
	}
}
func RunGates(text string, gates []Gate) []Finding {
	o := []Finding{}
	for _, g := range gates {
		if !g.Check(text) {
			sev := 1
			if g.Required {
				sev = 3
			}
			o = append(o, Finding{"gate." + g.Name, g.Name + " gate failed", sev})
		}
	}
	return o
}
func Passed(text string) bool { return Ready(RunGates(text, DefaultGates())) }
