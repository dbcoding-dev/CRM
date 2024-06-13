package auth

import (
	"goyave.dev/goyave/v5"
)

// RegisterController Register a new user
func RegisterController(response *goyave.Response, request *goyave.Request) {
	// Define the validation rules

	// Create the user
	user := User{
		Username: request.JSON["username"].(string),
		Email:    request.JSON["email"].(string),
		Password: request.JSON["password"].(string),
	}

	// Save the user
	if err := user.Save(); err != nil {
		response.JSON(500, err)
		return
	}

	response.JSON(201, user)
}
