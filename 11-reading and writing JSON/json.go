package main

import (
	"fmt"

	"github.com/oriastanjung/go-learn/helpers"
)

func main() {
	var user helpers.User = helpers.User{FirstName: "Orias"}
	fmt.Println(user)
}
