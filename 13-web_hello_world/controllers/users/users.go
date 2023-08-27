package users

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Create dummy user data

	users := []User{
		{ID: 1, Username: "user1", Email: "user1@example.com"},
		{ID: 2, Username: "user2", Email: "user2@example.com"},
		// Add more dummy users as needed
	}
	response := Response{
		Data:    users,
		Message: "",
	}
	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	if len(users) == 0 {
		response.Data = ""
		response.Message = "Not Found"
		responseJson, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		w.Write(responseJson)
		return
	}

	response.Data = users
	response.Message = "success"
	responseJson, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}
