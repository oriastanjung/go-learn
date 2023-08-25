package main

import "fmt"

// global variable = package level variable
var myName string

func main() {
	fmt.Println("variable")

	// buat variable
	// declaration
	// var namaVariable tipeData
	var whatToSay string = "Hello there"

	whatToSay = "hello orias"
	// dengan menggunakan :=
	// langsung mendeklare variable dan langsung beri value

	whatToSay1 := "hello ate"
	fmt.Println(whatToSay)
	fmt.Println(whatToSay1)

	var i int
	i = 10

	fmt.Println("i is ", i)

	whatToSaid := saySomething()

	fmt.Println(whatToSaid)
	var sum int = sumTwoNumbers(10, 6)

	fmt.Println("sum >> ", sum)

	checkThings1, checkThings2 := checkTwoThings()
	fmt.Println("checkThings1 >> ", checkThings1)
	fmt.Println("checkThings2 >> ", checkThings2)

}

// buat function
// func namaFunction() tipeDataReturnnya

func saySomething() string {
	return "Something"
}

func sumTwoNumbers(a int, b int) int {
	return a + b
}

func checkTwoThings() (string, string) {
	return "things1", "things2"
}
