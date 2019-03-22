package controllers

import (
	"encoding/json"
	"net/http"

	models "github.com/filipjedrasik/crr-api/models"
	u "github.com/filipjedrasik/crr-api/utils"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Błędne żądanie"))
		return
	}

	resp := user.Create()
	u.Respond(w, resp)
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Błędne żądanie"))
		return
	}

	response := models.Login(user.Email, user.Password)
	u.Respond(w, response)
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	response := u.Message(true, "Odświeżono")
	userId := r.Context().Value("userId").(uint)
	response["user"] = models.GetUser(uint(userId))
	u.Respond(w, response)
}
