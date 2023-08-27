package main

import (
	"fmt"
	"net/http"

	"github.com/oriastanjung/13-web_hello_world/controllers/users"
)

func main() {
	// http.HandleFunc only using GET method jadi kalo mau method lain wajib pakai framework seperti mux, gorilla, echo atau gin
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		n, err := fmt.Fprintf(writer, "hello world!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Number of bytes written: %d\n", n)

	})

	// kita bisa pisahin pake controller beda juga

	http.HandleFunc("/users", users.GetAllUsers)

	_ = http.ListenAndServe(":8080", nil)
}
