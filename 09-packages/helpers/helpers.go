package helpers // sesuaikan dengan nama folder
import "fmt"

type User struct {
	Name string
	Age  int
}

func (user *User) SaysName() {
	fmt.Println("my name is", user.Name)
}
