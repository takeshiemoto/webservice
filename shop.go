package structure

type Shop struct {
	name string
}

func (s *Shop) setName(name string) {
	s.name = name
}