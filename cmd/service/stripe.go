package service

import (
	"finchat-api/cmd/entities"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/customer"
	"github.com/stripe/stripe-go/v72/paymentintent"
)

func CreateCustomer(request *entities.CreateCustomerRequest) (*stripe.Customer, error) {
	params := &stripe.CustomerParams{
		Email:         stripe.String(request.Email),
		PaymentMethod: stripe.String(request.StripeCreditCardToken),
	}

	return customer.New(params)
}

func CreatePayment(request *entities.CreatePaymentRequest) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Customer: stripe.String(request.StripeCustomerID),
		Amount:   stripe.Int64(int64(request.Amount)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes: []*string{
			stripe.String("card"),
		},
	}

	return paymentintent.New(params)
}

func Payments(customerId string) *paymentintent.Iter {
	params := &stripe.PaymentIntentListParams{
		Customer: stripe.String(customerId),
	}

	return paymentintent.List(params)
}
