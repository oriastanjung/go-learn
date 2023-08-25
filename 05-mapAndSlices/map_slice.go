package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// map itu mirip kyk struct cmn dia tipenya sama
	// namaVariable = make(map[tipeDataKey]tipeDataValue)
	myMap := make(map[string]string)
	myMap["dog"] = "cihuahua"
	myMap["cat"] = "anggora"

	log.Println(myMap["dog"])
	log.Println(myMap["cat"])

	// kita bisa reassign value juga
	myMap["dog"] = "husky"
	log.Println(myMap["dog"])

	myMap2 := make(map[string]int)
	myMap2["first"] = 1

	log.Println(myMap2["first"])

	// kita bisa kasih return sbuah struct juga
	myMap3 := make(map[string]User)
	me := User{
		FirstName: "Orias",
		LastName:  "tanjung",
	}
	myMap3["me"] = me

	log.Println(myMap3["me"].FirstName)
	// pada map bakalan diacak terus, sehingga aksesnya wajib menggunakan key tidak menggunakan index

	// di go lang, kita tidak mengenal array
	// tapi adanya slices, cara kerja nya mirip array

	// syntax slice
	// var namaVariable []tipeData
	var mySlice []string
	mySlice = append(mySlice, "Orias")
	mySlice = append(mySlice, "Oria1")
	log.Println(mySlice)
	log.Println(mySlice[0])

	mySlice2 := []int{3, 2, 5, 4, 1}
	log.Println(mySlice2)

	// sort.Slice(mySlice2, func(i, j int) bool {
	// 	return mySlice2[i] < mySlice2[j]
	// })

	sort.Ints(mySlice2)
	log.Println(mySlice2)
	log.Println(mySlice2[0:2]) // dari index 0 ke 2

}
