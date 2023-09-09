package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	ID        string
	Name      string
	Role      string
	Email     string
	Phone     string
	Contacted bool
}

var customers = []Customer{
	{
		ID:        "1",
		Name:      "Nam",
		Role:      "Software Engineer",
		Email:     "nam14nt@gmail.com",
		Phone:     "0353930854",
		Contacted: true,
	},
	{
		ID:        "2",
		Name:      "Nam",
		Role:      "Software Engineer",
		Email:     "nam14nt@gmail.com",
		Phone:     "0353930854",
		Contacted: true,
	},
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, customer := range customers {
		if customer.ID == params["id"] {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(customer)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Customer{})
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customer Customer
	_ = json.NewDecoder(r.Body).Decode(&customer)
	customers = append(customers, customer)
	json.NewEncoder(w).Encode(customer)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, customer := range customers {
		if customer.ID == params["id"] {
			customers = append(customers[:i], customers[i+1:]...)
			var updatedCustomer Customer
			_ = json.NewDecoder(r.Body).Decode(&updatedCustomer)
			updatedCustomer.ID = params["id"]
			customers = append(customers, updatedCustomer)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updatedCustomer)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(customers)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, customer := range customers {
		if customer.ID == params["id"] {
			customers = append(customers[:i], customers[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(customers)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	fmt.Println("Server is starting on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
