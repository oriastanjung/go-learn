package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// syntax
// biasanya dibuat di package level
// biar bisa diakes semuanya

// type namaStruct struct
type User struct {
	FirstName   string // alasan depannya smua kapital karena untuk public nntinya dapat akses smua data variable atau function
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

// type namaStruct struct
type UserJson struct {
	FirstName   string    `json:"firstName"` // alasan depannya smua kapital karena untuk public nntinya dapat akses smua data variable atau function
	LastName    string    `json:"lastName"`
	PhoneNumber string    `json:"phoneNumber"`
	Age         int       `json:"age"`
	BirthDate   time.Time `json:"birthDate"`
}

func main() {
	orias := User{
		FirstName:   "Orias",
		LastName:    "Tanjung",
		PhoneNumber: "082269727772",
		Age:         21,
	}

	var omi User = User{
		FirstName: "Omi",
		LastName:  "Diyanto",
	}

	log.Println("orias >> ", orias.BirthDate)
	log.Println("omi >> ", omi)

	arr := []User{
		omi,
		orias,
	}

	for index := range arr {
		log.Println(arr[index])
	}

	// testing json
	person1 := UserJson{
		FirstName:   "orias",
		LastName:    "tanjung",
		PhoneNumber: "+6282269282",
		Age:         21,
	}

	// encode struct ke json
	jsonData, err := json.Marshal(person1)
	if err != nil {
		log.Fatal(err)
	}

	data := string(jsonData)
	fmt.Println("data >> ", data)
}

func test() { // ini hanya bisa diakses di file ini saja karena awal huruf bukan kapital

}

func Test1() { // ini bisa diakses oleh file mana saja nantinya karena di awal huruf kapital

}
