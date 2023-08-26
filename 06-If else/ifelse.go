package main

import "log"

func main() {
	animal := "dog"

	if animal == "cat" {
		log.Println("Animal is cat")
	} else if animal == "dog" {
		log.Println("Animal is dog")
	} else {
		log.Println("it is animal")
	}

	myNum := 100
	isTrue := false

	if myNum > 99 && isTrue {
		log.Println("myNum is greater than 99 and isTrue is set to true")
	} else if myNum > 99 && !isTrue {
		log.Println("myNum is greater than 99 and isTrue is set to false")
	}

	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println("cat")
	case "dog":
		log.Println("dog")
	case "fish":
		log.Println("fish")
	default:
		log.Println("animal")
	}

}
