package main

import "fmt"

const (
	newspaperPubType = "newspaper"
	magazinePubType  = "magazine"
)

func newPublication(pubType, name, pub string, pages int) (iPublication, error) {
	if pubType == newspaperPubType {
		return createNewspaper(name, pub, pages), nil
	}
	if pubType == magazinePubType {
		return createNewMagazine(name, pub, pages), nil
	}
	return nil, fmt.Errorf("no such publication type")
}
