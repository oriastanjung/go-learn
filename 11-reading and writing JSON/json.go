package main

import (
	"encoding/json"
	"fmt"

	"github.com/oriastanjung/go-learn/helpers"
)

func main() {
	myJson := `
	[
		{
			"first_name" : "clark",
			"last_name": "kent",
			"age" : 21,
			"has_dog" : true
		},
		{
			"first_name" : "bruce",
			"last_name": "wayne",
			"age" : 23,
			"has_dog" : false
		}
	]
	`

	var user helpers.User = helpers.User{FirstName: "Orias"}
	fmt.Println(user)

	//reading json
	var unmarshalledVariable []helpers.Person
	err := json.Unmarshal([]byte(myJson), &unmarshalledVariable)
	if err != nil {
		fmt.Println("Error unmarshalling json", err)
	}
	fmt.Printf("unmarshalled %v", unmarshalledVariable)
	fmt.Printf("unmarshalled[0] %v", unmarshalledVariable[0])
	fmt.Printf("unmarshalled[1] %v", unmarshalledVariable[1])

	// write json
	var newData []helpers.Person
	var p1 helpers.Person
	p1.FirstName = "diana"
	p1.LastName = "prince"
	p1.Age = 100
	p1.HasDog = false

	var p2 helpers.Person
	p2.FirstName = "wally"
	p2.LastName = "west"
	p2.Age = 25
	p2.HasDog = false

	newData = append(newData, p1, p2)
	fmt.Println("newData before JSON >> ", newData)
	newjson, err := json.MarshalIndent(newData, "", "    ")
	if err != nil {
		fmt.Println("error marshalling", err)
	}
	fmt.Println("newJson before casting to string >>", newjson)
	fmt.Println("newJson after casting to string >>", string(newjson))
	fmt.Println("\n\ncombining the json before and after newjson")

	unmarshalledVariable = append(unmarshalledVariable, newData...)
	fmt.Println("\n\nnewData with JSON >> ", unmarshalledVariable)
	allData, err := json.MarshalIndent(unmarshalledVariable, "", "    ")
	if err != nil {
		fmt.Println("error marshalling ", err)
	}
	fmt.Println("newData with JSON >> ", string(allData))
	// pakai json.MarshallIndent() biar rapi aja di dev, nntinya kita juga pakai .Marshall() aja kok

}
