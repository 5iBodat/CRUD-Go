package router

import (
	"golang-rest-api/controller/departement"
	"golang-rest-api/controller/designation"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	// departement
	r.HandleFunc("/departement", departement.Read).Methods("GET")
	r.HandleFunc("/departement/create", departement.Create).Methods("POST")
	r.HandleFunc("/departement/update/{id}", departement.Update).Methods("PUT")    // Define parameter with {id}
	r.HandleFunc("/departement/delete/{id}", departement.Delete).Methods("DELETE") // Define parameter with {id}

	//designation
	r.HandleFunc("/designation", designation.Read).Methods("GET")
	r.HandleFunc("/designation/create", designation.Create).Methods("POST")
	r.HandleFunc("/designation/update/{id}", designation.Update).Methods("PUT")    // Define parameter with {id}
	r.HandleFunc("/designation/delete/{id}", designation.Delete).Methods("DELETE") // Define parameter with {id}

	return r
}
