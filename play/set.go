package play

import (
	"fmt"
	"slices"

)

type Set struct {
	items []string
}

func (s *Set) Add(item string) {
	if slices.Contains(s.items, item) {
		return
	}

	s.items = append(s.items, item)
}

func (s *Set) String() string {
	return fmt.Sprintf("%v", s.items)
}

