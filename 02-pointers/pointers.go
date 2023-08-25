package main

import (
	"log"
)

// global variable
var sName = "halo"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString >> ", myString)

	//pass the address
	chaneUsingPointer(&myString)
	log.Println("myString >> ", myString)
	log.Println("sName >> ", sName)
}

// coba buat tipe data pointer

// nama variable *tipepointer

func chaneUsingPointer(stringSaid *string) {
	newValue := "Red"
	*stringSaid = newValue
	log.Println("sName >> ", sName)

}
