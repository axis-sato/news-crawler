package entities

import "fmt"

type Tag struct {
	ID   int
	Name string
}

func (t Tag) String() string {
	return fmt.Sprintf("id: %v, name: %v", t.ID, t.Name)
}

type Tags []Tag

func (tags *Tags) Names() []string {
	var names []string
	for _, t := range *tags {
		names = append(names, t.Name)
	}
	return names
}
