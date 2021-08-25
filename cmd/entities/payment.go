package entities

type CreatePaymentRequest struct {
	StripeCustomerID string
	Amount           int
}

type CreatePaymentResponse struct {
	PaymentIntentID string `json:"paymentIntentID"`
}

type Payment struct {
	ID     string `json:"id"`
	Amount int64  `json:"amount"`
}

type PaymentsResponse struct {
	Payments []Payment `json:"payments"`
}
