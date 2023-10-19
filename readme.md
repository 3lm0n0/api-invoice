```
package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

type InvoiceController struct {}

func (ic *InvoiceController) Get(w http.ResponseWriter, r *http.Request) {}

func (ic *InvoiceController) Post(w http.ResponseWriter, r *http.Request) {}

func (ic *InvoiceController) Delete(w http.ResponseWriter, r *http.Request) {}

func main() {
	router := mux.NewRouter()
	
	ic := &InvoiceController{}

	router.HandleFunc("/invoice/{id:[0-9]+}", ic.Get).Methods("GET")

	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
}
```

Crear factura
```
{
    "lines": [
        {
            "product": "Foo Bar",
            "id_product": 1,
            "quantity": 0,
            "price": 0,
            "tax_rate": 0,
            "discount_rate": 0,
            "currency": "USD"
        }
    ],
    "client": {
        "name": "Julio Health",
        "identification": "1-1111-1111"
    }
}
```

Obtener factura
```
{
    "response": [
        {
            "id": 0,
            "lines": [
                {
                    "id": 0,
                    "quantity": 0,
                    "price": 0,
                    "price_crc": 0,
                    "tax_rate": 0,
                    "discount_rate": 0,
                    "currency": "USD"
                }
            ],
            "client": {
                "name": "Julio Health",
                "identification": "1-1111-1111"
            },
            "tax_total": 0,
            "discount_total": 0,
            "subtotal": 0,
            "total": 0,
            "payments": [
                {
                    "id": "",
                    "total": ""
                }
            ],
            "balance": 0
        }
    ]
} 
```
```
