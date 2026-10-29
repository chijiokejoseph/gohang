package settings

import "log"

type Mode int

const (
	Beginner Mode = iota
	Amateur
	Regular
	Professional
)

func (m Mode) String() string {
	switch m {
	case Beginner:
		return "Beginner"
	case Amateur:
		return "Amateur"
	case Regular:
		return "Regular"
	case Professional:
		return "Professional"
	default:
		log.Fatal("invalid Mode specified")
	}
	return ""
}
