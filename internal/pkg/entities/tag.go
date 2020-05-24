package entities

import "fmt"

type Tag struct {
	ID   int
	Name string
}

func NewTag(id int, name string) *Tag {
	return &Tag{
		ID:   id,
		Name: name,
	}
}

func (t Tag) String() string {
	return fmt.Sprintf("{id: %v, name: %v}", t.ID, t.Name)
}

type Tags []Tag

func NewTags(ts []Tag) *Tags {
	var tags Tags
	for _, t := range ts {
		tags = append(tags, t)
	}
	return &tags
}

func (tags *Tags) Names() []string {
	var names []string
	for _, t := range *tags {
		names = append(names, t.Name)
	}
	return names
}
