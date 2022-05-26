package models

import "github.com/xendit/xendit-go"

type ICustomer struct {
	GivenNames   string               `json:"given_names"`
	Email        string               `json:"email"`
	MobileNumber string               `json:"mobile_number"`
	Address      string               `json:"address"`
	Items        []xendit.InvoiceItem `json:"items"`
}

type EWalletModel struct {
	Price  float64 `json:"price"`
	Method string  `json:"method"`
	Phone  string  `json:"phone"`
}
