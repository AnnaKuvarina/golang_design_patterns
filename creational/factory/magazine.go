package main

import "fmt"

type magazine struct {
	publication
}

func (m magazine) String() string {
	return fmt.Sprintf("This is a magazine named: %s", m.name)
}

func createNewMagazine(name, publisher string, pages int) iPublication {
	return &magazine{
		publication: publication{
			name:      name,
			pages:     pages,
			publisher: publisher,
		},
	}
}
