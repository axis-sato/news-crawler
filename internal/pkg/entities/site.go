package entities

import (
	"fmt"
)

type Site struct {
	ID   int
	Name string
	URL  string
}

func (s Site) String() string {
	return fmt.Sprintf("{id: %v, name: %v, url: %v}", s.ID, s.Name, s.URL)
}
