package Controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"log"
	"myapp/Models"
	"net/http"
	"os"
)

func HandleCreatePaymentIntent(c echo.Context) error {
	stripe.Key = os.Getenv("STRIPE_KEY")

	payment := new(Models.Payment)

	fmt.Printf("Add new payment \n")

	if err := c.Bind(payment); err != nil {
		return c.String(http.StatusBadRequest, "Bad payment "+err.Error())
	}

	//Database.Database.Create(payment)

	// Create a PaymentIntent with amount and currency
	params := &stripe.PaymentIntentParams{
		Amount:      stripe.Int64(int64(payment.Amount * 100)),
		Currency:    stripe.String(string(stripe.CurrencyPLN)),
		Description: stripe.String("Ruczaj restaurant"),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	pi, err := paymentintent.New(params)
	log.Printf("pi.New: %v", pi.ClientSecret)

	if err != nil {
		log.Printf("pi.New: %v", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, pi.ClientSecret)
}
