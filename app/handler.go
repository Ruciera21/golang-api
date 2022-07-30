package app

import (
	"encoding/json"
	"fmt"
	"goapi-nunu/service"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	fmt.Println("status =>", status)
	cust, err := ch.service.GetAllCustomer(status)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, cust)
}

func (ch *CustomerHandler) getCustomerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["customer_id"]
	customer, err := ch.service.GetCustomerByID(customerID)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, customer)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	// set header terlebih dahulu, baru writeheadernya, jika terbalik maka akan lock header dengan seeting text/plain
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

// func addCustomer(w http.ResponseWriter, r *http.Request) {
// 	// * decode request body
// 	var cust Customer
// 	json.NewDecoder(r.Body).Decode(&cust)

// 	// * generate new id
// 	nextID := getNextID()
// 	cust.ID = nextID

// 	// * save data to array
// 	customers = append(customers, cust)

// 	w.WriteHeader(http.StatusCreated)
// 	fmt.Fprintln(w, "customer successfully created")
// }

// func updateCustomer(w http.ResponseWriter, r *http.Request) {

// 	// * get route variable
// 	vars := mux.Vars(r)

// 	customerId := vars["customer_id"]

// 	// * convert string to int
// 	id, err := strconv.Atoi(customerId)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprint(w, "invalid customer id")
// 		return
// 	}

// 	// * searching customer data
// 	var cust Customer

// 	for customerIndex, data := range customers {
// 		if data.ID == id {
// 			// * save temp data for validation
// 			cust = data

// 			// * decode request body
// 			var newCust Customer
// 			json.NewDecoder(r.Body).Decode(&newCust)

// 			// * do update
// 			customers[customerIndex].Name = newCust.Name
// 			customers[customerIndex].City = newCust.City
// 			customers[customerIndex].ZipCode = newCust.ZipCode

// 			// fmt.Println(customers)

// 			w.WriteHeader(http.StatusOK)
// 			fmt.Fprintln(w, "customer data updated")
// 			return
// 		}
// 	}

// 	// ! error if data not find
// 	if cust.ID == 0 {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(w, "customer data not found")
// 		return
// 	}

// }

// func getNextID() int {
// 	lastIndex := len(customers) - 1
// 	lastCustomer := customers[lastIndex]
// 	// cust := customers[len(customers)-1]

// 	return lastCustomer.ID + 1
// }
