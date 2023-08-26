package main

import (
	"log"
	"math/rand"
	"time"
)

func RandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	// syntax buat channel
	// variable := make(chan tipeChannel)
	testChannelInt := make(chan int)
	defer close(testChannelInt) // setiap buat channel sangat kritis untuk close channel nya setelah program selesai

	go CalculateValue(testChannelInt) // baru kita run channel dengan keyword go
	num := <-testChannelInt           // artinya variable num akan mengambil nilai dari channel testChannelInt
	log.Println(num)

}
