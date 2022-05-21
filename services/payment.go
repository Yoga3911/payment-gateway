package services

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/balance"
	"github.com/xendit/xendit-go/ewallet"
	"github.com/xendit/xendit-go/invoice"
)

type Payment interface {
	GetBalance(c *fiber.Ctx) error
	EWalletCharge(c *fiber.Ctx) error
	GetEWalletCharge(c *fiber.Ctx) error
	CreateInvoice(c *fiber.Ctx) error
}

type payment struct {
	secret string
}

func NewPayment(secret string) Payment {
	return &payment{
		secret: secret,
	}
}

func (p *payment) GetBalance(c *fiber.Ctx) error {
	xendit.Opt.SecretKey = p.secret

	data := balance.GetParams{
		AccountType: "CASH",
	}

	resp, err := balance.Get(&data)
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "Get balance success",
		"data":    resp.Balance,
	})
}

func (p *payment) EWalletCharge(c *fiber.Ctx) error {
	xendit.Opt.SecretKey = p.secret

	data := ewallet.CreateEWalletChargeParams{
		ReferenceID:    "test-reference-id",
		Currency:       "IDR",
		Amount:         1000,
		CheckoutMethod: "ONE_TIME_PAYMENT",
		ChannelCode:    "ID_SHOPEEPAY",
		ChannelProperties: map[string]string{
			"success_redirect_url": "https://dashboard.xendit.co/register/1",
		},
		Metadata: map[string]interface{}{
			"branch_code": "tree_branch",
		},
	}

	charge, chargeErr := ewallet.CreateEWalletCharge(&data)
	if chargeErr != nil {
		log.Fatal(chargeErr)
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "Create new e-wallet charge success",
		"data":    charge,
	})
}

func (p *payment) GetEWalletCharge(c *fiber.Ctx) error {
	xendit.Opt.SecretKey = p.secret

	data := ewallet.GetEWalletChargeStatusParams{
		ChargeID: "ewc_27a2f5a4-93c0-48f5-a7a8-24fcb3ff89ee",
	}

	charge, chargeErr := ewallet.GetEWalletChargeStatus(&data)
	if chargeErr != nil {
		log.Fatal(chargeErr)
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "Get e-wallet charge success",
		"data":    charge,
	})
}

func (p *payment) CreateInvoice(c *fiber.Ctx) error {
	xendit.Opt.SecretKey = p.secret

	customer := xendit.InvoiceCustomer{
		GivenNames:   "Kevin",
		Email:        "kevin@example.com",
		MobileNumber: "+6287774441111",
		Address:      "Jalan Kenanga 12",
	}

	items := []xendit.InvoiceItem{
		{
			Name:     "Air Conditioner",
			Quantity: 4,
			Price:    100000,
			Category: "Electronic",
			Url:      "https://yourcompany.com/example_item",
		},
	}

	fees := []xendit.InvoiceFee{
		{
			Type:  "ADMIN",
			Value: 5000,
		},
	}

	NotificationType := [3]string{"whatsapp", "email", "sms"}

	customerNotificationPreference := xendit.InvoiceCustomerNotificationPreference{
		InvoiceCreated:  NotificationType[:],
		InvoiceReminder: NotificationType[:],
		InvoicePaid:     NotificationType[:],
		InvoiceExpired:  NotificationType[:],
	}

	data := invoice.CreateParams{
		ExternalID:                     "demo_1475801962608",
		Amount:                         50000,
		Description:                    "Invoice Demo #1234",
		InvoiceDuration:                86400,
		Customer:                       customer,
		CustomerNotificationPreference: customerNotificationPreference,
		Currency:                       "IDR",
		Items:                          items,
		Fees:                           fees,
		
	}

	resp, err := invoice.Create(&data)
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": "Create invoice success",
		"data":    resp,
	})
}
