package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type InvoiceController struct {
	is InvoiceServiceInterface
}

func NewInvoiceController(is InvoiceServiceInterface) *InvoiceController {
	return &InvoiceController{is: is}
}

func (ic *InvoiceController) Get(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		fmt.Println("ID conversion failed", err)
		panic(err)
	}

	i, err := ic.is.GetInvoice(id)
	if err != nil {
		fmt.Println("Invoice not found", err)
		panic(err)
	}

	// console check
	fmt.Println(i)
	
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json; charset=UTF-8")
	w.Write([]byte(fmt.Sprintf("Invoice %v", i)))
}

//func (ic *InvoiceController) Post(w http.ResponseWriter, r *http.Request) {}

//func (ic *InvoiceController) Delete(w http.ResponseWriter, r *http.Request) {}

func main() {
	router := mux.NewRouter()
	
	is := NewInvoiceService() // New invoice service interface instance
	ic := NewInvoiceController(is) // Pass the invoice service interface to the controller

	router.HandleFunc("/invoice/{id:[0-9]+}", ic.Get).Methods("GET")

	http.Handle("/", router)

	http.ListenAndServe(":8080", router) // runs the server
}
