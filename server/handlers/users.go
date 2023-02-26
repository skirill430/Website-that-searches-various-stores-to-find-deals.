package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/skirill430/Quick-Shop/server/utils/db"
	"gorm.io/gorm"
)

/*
CreateUser Sample request body (json):

	{
		"username": "Daniel",
		"password": "password"
	}
*/
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user db.User
	json.NewDecoder(r.Body).Decode(&user)

	// check credentials are valid
	if user.Username == "" || user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Request Body Missing Fields"))
		return
	}

	// add user only if username doesn't exist in database already
	res := db.DB.Where("username = ?", user.Username).FirstOrCreate(&user)

	if res.RowsAffected == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Sorry, this username is already taken. Enter another username."))
		return
	}
}

/*
CreateUser Sample request body (json):

	{
		"username": "admin",
		"password": "123456"
	}
*/
func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type Credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var inputCredentials Credentials
	json.NewDecoder(r.Body).Decode(&inputCredentials)

	// check credentials are valid
	if inputCredentials.Username == "" || inputCredentials.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Request Body Missing Fields"))
		return
	}

	var dbUser *db.User
	// cannot find username in database
	err := db.DB.First(&dbUser, "username = ?", inputCredentials.Username).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("This username does not exist. Create an account."))
		return
	}

	// password does not match
	if inputCredentials.Password != dbUser.Password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Incorrect password. Try again."))
		return
	}

	json.NewEncoder(w).Encode(dbUser)
}
