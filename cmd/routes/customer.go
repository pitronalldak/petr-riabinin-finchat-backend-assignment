package routes

import (
	"finchat-api/cmd/entities"
	"finchat-api/cmd/service"

	"github.com/gofiber/fiber/v2"
)

func CustomerRouter(app fiber.Router) {
	app.Post("/customer", createCustomer)

}

func createCustomer(c *fiber.Ctx) error {
	var requestBody entities.CreateCustomerRequest
	err := c.BodyParser(&requestBody)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	result, err := service.CreateCustomer(&requestBody)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.JSON(entities.CreateCustomerResponse{StripeCustomerID: result.ID})
}
