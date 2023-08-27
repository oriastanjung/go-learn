package main

import (
	"errors"
	"fmt"
)

// file test harus benar yaitu namaFile_test.go harus ada _test nya untuk spesific file yang mau ditest

func main() {
	result, err := divide(100.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func divide(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("cannot divide with 0")
	}
	result = x / y
	return result, nil
}
