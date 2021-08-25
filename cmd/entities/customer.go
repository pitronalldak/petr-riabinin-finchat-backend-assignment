package entities

type CreateCustomerRequest struct {
	Email                 string
	StripeCreditCardToken string
}

type CreateCustomerResponse struct {
	StripeCustomerID string `json:"stripeCustomerID"`
}
