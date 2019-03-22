package main

import (
	"fmt"
	"net/http"
	"os"

	app "github.com/filipjedrasik/crr-api/app"
	controllers "github.com/filipjedrasik/crr-api/controllers"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println("It works on http://localhost:" + port + "/api address")

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})

	err := http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router))
	if err != nil {
		fmt.Print(err)
	} 

}
