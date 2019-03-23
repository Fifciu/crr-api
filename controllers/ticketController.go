package controllers

import (
	"net/http"

	models "github.com/filipjedrasik/crr-api/models"
	u "github.com/filipjedrasik/crr-api/utils"
)

func Ticket(w http.ResponseWriter, r *http.Request) {
	response := u.Message(true, "Zakupiono bilet")
	userId := r.Context().Value("userId").(uint)
	response["ticketExpires"] = models.BuyTicket(userId)
	u.Respond(w, response)
}
