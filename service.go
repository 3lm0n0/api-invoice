package main

import (
	m "facturas/models"
)

type InvoiceServiceInterface interface {
	GetInvoice(id int) (*m.Invoice, error)
}

type InvoiceService struct {
}


func NewInvoiceService() InvoiceServiceInterface {
	return &InvoiceService{}
}

// GetInvoice implements InvoiceServiceInterface.
func (*InvoiceService) GetInvoice(id int) (*m.Invoice, error) {

	i := m.Invoice{
		Id_product:    123,
		Product:       "test",
		Quantity:      10,
		Price:         1,
		Tax_rate:      1,
		Discount_rate: 0,
		Currency:      "USD",
	}

	return &i, nil
}
