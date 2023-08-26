package main

import "fmt"

// interfacce merupakan type untuk kumpulan method yang sama yang dimiliki oleh beberapa type struct lainnya

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name string
	Type string
}

type Bird struct {
	Name  string
	Color string
}

func main() {
	dog := Dog{
		Name: "Chika",
		Type: "Chihuahua",
	}

	bird := Bird{
		Name:  "Kakao",
		Color: "Green",
	}

	PrintInfo(&dog)
	fmt.Println()
	PrintInfo(&bird)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has a", a.NumberOfLegs(), "legs")
}

// menggunakan receivers
func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}
func (d *Bird) Says() string {
	return "kiwkiw"
}

func (d *Bird) NumberOfLegs() int {
	return 2
}
