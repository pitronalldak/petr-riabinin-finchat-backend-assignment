package routes

import (
	"finchat-api/cmd/entities"
	"finchat-api/cmd/service"

	"github.com/gofiber/fiber/v2"
)

func PaymentRouter(app fiber.Router) {
	app.Post("/payments", createPayment)
	app.Get("/payments/:customer_id", getPayments)

}

func createPayment(c *fiber.Ctx) error {
	var requestBody entities.CreatePaymentRequest
	err := c.BodyParser(&requestBody)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	result, err := service.CreatePayment(&requestBody)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.JSON(entities.CreatePaymentResponse{PaymentIntentID: result.ID})
}

func getPayments(c *fiber.Ctx) error {
	customerId := c.Params("customer_id")
	iter := service.Payments(customerId)

	var payments []entities.Payment

	for iter.Next() {
		item := iter.PaymentIntent()
		payments = append(payments, entities.Payment{ID: item.ID, Amount: item.Amount})
	}

	if err := iter.Err(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.JSON(entities.PaymentsResponse{Payments: payments})
}
