package main

import (
	"fmt"
	"net/http"

	"github.com/oriastanjung/bookings-udemy-21/controllers/users"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		n, err := fmt.Fprintf(writer, "hello world!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Number of bytes written: %d\n", n)

	})

	http.HandleFunc("/users", users.GetAllUsers)

	_ = http.ListenAndServe(":8080", nil)
}
