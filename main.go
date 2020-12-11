package main

import (
	"address-list/controllers"
	"address-list/driver"
	"address-list/models"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var address_users []models.Address
var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db = driver.ConnectDB()
	controller := controllers.Controllers{}
	router := mux.NewRouter()

	router.HandleFunc("/address", controller.GetAddressUsers(db)).Methods("GET")
	router.HandleFunc("/address/{id}", controller.GetAddress(db)).Methods("GET")
	router.HandleFunc("/address/users/{id}", controller.GetInfoAddress(db)).Methods("GET")
	router.HandleFunc("/address", controller.AddAddressUser(db)).Methods("POST")
	router.HandleFunc("/address", controller.UpdateAddressUser(db)).Methods("PUT")
	router.HandleFunc("/address/{id}", controller.RemoveAddressUser(db)).Methods("DELETE")

	fmt.Println("Server is running at port 8001")
	log.Fatal(http.ListenAndServe(":8001", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(router)))
}
