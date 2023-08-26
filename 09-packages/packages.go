package main

// package ada disini yang dari import
import (
	"fmt"

	"github.com/oriastanjung/go-learn/09-packages/helpers"
)

func main() {
	fmt.Println("Hello")
	var user helpers.User = helpers.User{
		Name: "orias",
		Age:  21,
	}

	fmt.Println(user)
	user.SaysName()
}

// jika di js ada npm di go kita pakainya go mod
// inisialisasi awal dengan

// go mod init github.com/namaUser/namaRepo
// pada terminal

// untuk penamaan lokasi sebenarnya lebih ke konvensional meskipun nantinya tidak di push ke github
