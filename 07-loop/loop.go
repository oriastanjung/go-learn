package main

import (
	"fmt"
	"log"
)

func main() {
	// syntax
	// for deklarasi; kondisiCheck; incrementOrDecrement {}
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	log.Println()

	animals := []string{
		"dog", "fish", "horse", "cow", "cat",
	}

	// for indexSekarang, ValuePadaIndex := range KumpulanData
	for i, animal := range animals {
		log.Println(i, animal) // i kita padahal ga butuh, kita bisa ganti dengan _ biar ga kepakai
	}

	log.Println()

	// for _, ValuePadaIndex := range KumpulanData
	for _, animal := range animals {
		log.Println(animal)
	}

	log.Println()

	// dengan tipe map kita akan looping juga, hanya saja index nya menjadi key nya map

	birds := make(map[string]string)
	birds["parrot"] = "Kakao"
	birds["eagle"] = "Mighty"

	// for key, ValuePadaKey := range map
	for birdsTypeKey, bird := range birds {
		log.Println(birdsTypeKey, bird)
	}

	log.Println()

	// bisa looping dengan string juga

	firstLine := "Once Upon a time"
	for i, letter := range firstLine {
		log.Println(i, ":", string(letter))
	}
	log.Println()

	// loop sangat berguna juga untuk print data models seperti struct

	type User struct {
		FullName string
		Age      int
		Email    string
	}

	var users []User
	users = append(users, User{FullName: "Orias", Age: 21, Email: "Oriastan999@gmail.com"})
	users = append(users, User{FullName: "Omi", Age: 31, Email: "omidiyanto9@gmail.com"})
	users = append(users, User{FullName: "Seto", Age: 16, Email: "setograph@gmail.com"})

	fmt.Println("FullName\t\tEmail\t\t\tAge")
	for _, user := range users {
		fmt.Println(user.FullName, "\t\t", user.Email, "\t\t", user.Age)
	}
}
