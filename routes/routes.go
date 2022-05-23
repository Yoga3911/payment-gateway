package routes

import (
	"os"
	"payment/services"

	"github.com/gofiber/fiber/v2"
)

func Data(app *fiber.App) {
	payment := services.NewPayment(os.Getenv("KEY"))
	api := app.Group("/api/v1")
	api.Get("/balance", payment.GetBalance)
	api.Post("/ewallet", payment.EWalletCharge)
	api.Get("/ewallet", payment.GetEWalletCharge)
	api.Post("/invoice", payment.CreateInvoice)
	api.Post("/customer", payment.CreateCustomer)
}