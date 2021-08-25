package main

import (
	"log"
	"os"

	"finchat-api/cmd/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/stripe/stripe-go/v72"
)

func main() {

	port := ":8080"
	sk_test_your_key := os.Getenv("STRIPE_SK")

	stripe.Key = sk_test_your_key

	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")
	routes.CustomerRouter(api)
	routes.PaymentRouter(api)

	log.Fatal(app.Listen(port))
}
