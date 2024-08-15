package main

import "fmt"

func main() {
	mag1, _ := newPublication(magazinePubType, "Time", "The Times", 50)
	mag2, _ := newPublication(magazinePubType, "Life", "Inc Life", 40)
	news1, _ := newPublication(newspaperPubType, "The Herald", "Heralders", 60)
	news2, _ := newPublication(newspaperPubType, "The Standard", "Standarders", 30)

	printPubDetails(mag1)
	printPubDetails(mag2)
	printPubDetails(news1)
	printPubDetails(news2)
}

func printPubDetails(pub iPublication) {
	fmt.Println("-----------")
	fmt.Printf("%s \n", pub)
	fmt.Printf("Type: %T\n", pub)
	fmt.Printf("Name: %s\n", pub.getName())
	fmt.Printf("Pages: %d\n", pub.getPages())
	fmt.Printf("Publisher: %s\n", pub.getPublisher())
	fmt.Println("-----------")
}
