package entity

import "time"

type Payment struct {
	ID_Payment          int       `json:"id_payment"`
	Payment_Code        string    `json:"payment_code"`
	Payment_Merchant    string    `json:"payment_destination"`
	Payment_Amount      float32   `json:"payment_amount"`
	Payment_Description string    `json:"payment_description"`
	Date_Time           time.Time `json:"date_time"`
}

type PaymentRequest struct {
	Payment_Code        string  `json:"payment_code"`
	Payment_Merchant    string  `json:"payment_destination"`
	Payment_Amount      float32 `json:"payment_amount"`
	Payment_Description string  `json:"payment_description"`
}
