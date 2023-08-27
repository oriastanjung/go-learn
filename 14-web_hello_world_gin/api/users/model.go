package users

type User struct {
	Id       string `json:"_id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
