package controllers

import (
	"address-list/models"
	addressRepository "address-list/repository/address"
	"address-list/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Controllers struct{}

var address_users []models.Address

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controllers) GetAddressUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var address models.Address
		var error models.Error

		address_users = []models.Address{}
		addressRepo := addressRepository.AddressRepository{}
		address_users, err := addressRepo.GetAddressUsers(db, address, address_users)

		if err != nil {
			error.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, address_users)
	}
}

func (c Controllers) GetAddress(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var address models.Address
		var error models.Error

		params := mux.Vars(r)

		address_users = []models.Address{}
		addressRepo := addressRepository.AddressRepository{}

		id, _ := strconv.Atoi(params["id"])
		address, err := addressRepo.GetAddressUser(db, address, id)

		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Not Found"
				utils.SendError(w, http.StatusNotFound, error)
				return
			} else {
				error.Message = "Server Error"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, address)
	}
}

func (c Controllers) AddAddressUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var address models.Address
		var addressID int
		var error models.Error

		json.NewDecoder(r.Body).Decode(&address)

		if address.Street == "" || address.City == "" || address.Zip == "" || address.UserID == 0 {
			error.Message = "Enter missing fields."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		addressRepo := addressRepository.AddressRepository{}
		addressID, err := addressRepo.AddAddressUser(db, address)

		if err != nil {
			error.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "text/plan")
		utils.SendSuccess(w, addressID)
	}
}

func (c Controllers) UpdateAddressUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var address models.Address
		var error models.Error

		json.NewDecoder(r.Body).Decode(&address)

		if address.ID == 0 || address.Street == "" || address.City == "" || address.Zip == "" {
			error.Message = "All fields are required"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		addressRepo := addressRepository.AddressRepository{}
		rowsUpdated, err := addressRepo.UpdateAddressUser(db, address)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsUpdated)
	}
}

func (c Controllers) RemoveAddressUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var error models.Error
		params := mux.Vars(r)
		addressRepo := addressRepository.AddressRepository{}
		id, _ := strconv.Atoi(params["id"])

		rowsDeleted, err := addressRepo.RemoveAddressUser(db, id)

		if err != nil {
			error.Message = "Server error."
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		if rowsDeleted == 0 {
			error.Message = "Not Found"
			utils.SendError(w, http.StatusNotFound, error)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsDeleted)
	}
}

func (c Controllers) GetInfoAddress(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var address models.Address
		var error models.Error

		params := mux.Vars(r)

		address_users = []models.Address{}
		addressRepo := addressRepository.AddressRepository{}

		id, _ := strconv.Atoi(params["id"])
		address, err := addressRepo.GetInfoAddress(db, address, id)
		fmt.Println(id)

		log.Print(address)

		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Not Found"
				utils.SendError(w, http.StatusNotFound, error)
				return
			} else {
				error.Message = "Server Error"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, address)
	}
}
