package helpers

type User struct {
	FirstName string
}

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	HasDog    bool   `json:"has_dog"`
}
