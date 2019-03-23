package main

import (
	"fmt"
	"net/http"
	"os"

	app "github.com/filipjedrasik/crr-api/go/app"
	controllers "github.com/filipjedrasik/crr-api/go/controllers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)

	router.HandleFunc("/api/user/register",
		controllers.CreateAccount).Methods("POST")

	router.HandleFunc("/api/user/login",
		controllers.Authenticate).Methods("POST")

	router.HandleFunc("/api/user/refresh",
		controllers.Refresh).Methods("GET")

	// ws://localhost:8000/api/chat?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjR9.sz7T_Jlg7jCC6ogiBHmZMUAVXn6rTkEaA9F3TVEh5u8
	router.HandleFunc("/api/chat/live",
		controllers.HandleConnection)

	router.HandleFunc("/api/chat/history",
		controllers.History)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Run WebSocket
	go controllers.HandleMessages()

	fmt.Println("It works on http://localhost:" + port + "/api address")

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})

	err := http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router))
	if err != nil {
		fmt.Print(err)
	}

}
