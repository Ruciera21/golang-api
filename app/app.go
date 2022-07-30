package app

import (
	"fmt"
	"goapi-nunu/domain"
	"goapi-nunu/logs"
	"goapi-nunu/service"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
	}
	for _, envKey := range envProps {
		if os.Getenv(envKey) == "" {
			logs.Fatal(fmt.Sprintf("environment variable %s not defined. Terminating application...", envKey))
		}
	}
}

func Start() {

	err := godotenv.Load()
	if err != nil {
		logs.Fatal("error loading .env file")
	}

	sanityCheck()

	// * wiring
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	// * create ServeMux
	mux := mux.NewRouter()

	// *defining routes

	mux.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	mux.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomerByID).Methods(http.MethodGet)

	// * starting server
	// http.ListenAndServe(":8080", mux)
	serverAddr := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")

	logs.Info(fmt.Sprintf("Starting server on %s:%s...", serverAddr, serverPort))
	http.ListenAndServe(fmt.Sprintf("%s:%s", serverAddr, serverPort), mux)
}
