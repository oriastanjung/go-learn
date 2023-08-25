package main

import "log"

type myStruct struct {
	FirstName string
}

// kita juga bisa kasih function ke struct dengan receivers
// syntax

// func (objekStruct *structNya) namaFunction() tipeData{}

func (obj *myStruct) printFirstName(greet string) string {
	return greet + " " + obj.FirstName
}

func main() {

	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Steph",
	}

	log.Println("myVar >> ", myVar.FirstName)
	log.Println("myVar2 >> ", myVar2.FirstName)
	// cara menggunakan receiver

	log.Println("myVar >> ", myVar.printFirstName("hello"))
}
