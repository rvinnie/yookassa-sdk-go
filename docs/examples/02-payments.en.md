## Working with payments

The SDK allows you to create, confirm, cancel payments, as well as receive information about them.

* [Request to create a payment](#Request-to-create-a-payment)
* [Request for payment confirmation](#Request-for-payment-confirmation)
* [Request to cancel an incomplete payment](#Request-to-cancel-an-incomplete-payment)
* [Get payment information](#Get-payment-information)
* [Get a list of payments with filtering](#Get-a-list-of-payments-with-filtering)

---

### Request to create a payment
[Creating a payment in documentation](https://yookassa.ru/developers/api?lang=en#create_payment)

To accept a payment, you need to create a payment object, `Payment`.
It contains all the necessary payment information (amount, currency, and status).
Payments have a linear life cycle, going from one status to the next sequentially.

The response to the request will contain the `Payment` object with the current status.

```go
import (
    "github.com/rvinnie/yookassa-sdk-go/yookassa"
    "github.com/rvinnie/yookassa-sdk-go/yookassa/common"
    "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
)

func main() {
	// Create a yookassa client by specifying the store ID and secret key
	yooclient := yookassa.NewClient('<Store ID>', '<Secret key>')
	// Create a payment handler
	paymentHandler := yookassa.NewPaymentHandler(yooclient)
	// Create a payment
	payment, _ := paymentHandler.CreatePayment(&yoopayment.Payment{
		Amount: &yoocommon.Amount{
			Value:    "1000.00",
			Currency: "RUB",
		},
		PaymentMethod: yoopayment.PaymentMethodType("bank_card"),
		Confirmation: yoopayment.Redirect{
			Type:      "redirect",
			ReturnURL: "https://www.example.com",
		},
		Description: "Test payment",
	})
}
```

---

### Request for payment confirmation

[Payment confirmation in documentation](https://yookassa.ru/developers/api?lang=en#capture_payment)

Confirm you’re ready to accept the payment. Once the payment is captured, the status will change to `succeeded`. 
After that, you can provide the customer with the product or service.

You can only capture payments with the `waiting_for_capture` status,
and only for a certain amount of time (depending on the payment method).
If you do not capture the payment within the allotted time, the status will change to `canceled`,
and the money will be returned to the user.

The response to the request will contain the `Payment` object with the current status.

```go
import (
    "github.com/rvinnie/yookassa-sdk-go/yookassa"
    "github.com/rvinnie/yookassa-sdk-go/yookassa/common"
    "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
)

func main() {
	// Create a yookassa client by specifying the store ID and secret key
	yooclient := yookassa.NewClient('<Store ID>', '<Secret key>')
	// Create a payment handler
	paymentHandler := yookassa.NewPaymentHandler(yooclient)
	// Create a payment
	payment, _ := paymentHandler.CreatePayment(&yoopayment.Payment{
		Amount: &yoocommon.Amount{
			Value:    "1000.00",
			Currency: "RUB",
		},
		PaymentMethod: yoopayment.PaymentMethodType("bank_card"),
		Confirmation: yoopayment.Redirect{
			Type:      "redirect",
			ReturnURL: "https://www.example.com",
		},
		Description: "Test payment",
	})

	// We are waiting for the payment to be made 30 seconds
	time.Sleep(time.Second * 30)
	
	// We confirm the payment
	payment, _ = paymentHandler.CapturePayment(payment)
}
```
[Learn more about confirming and canceling payments](https://yookassa.ru/developers/payment-acceptance/getting-started/payment-process?lang=en#capture-and-cancel)


---

### Request to cancel an incomplete payment
[Cancellation of payment in documentation](https://yookassa.ru/developers/api?lang=en#cancel_payment)

Cancel payments with the `waiting_for_capture` status.
Payment cancelation means you are not ready to dispatch a product or to provide a service to the user.
Once you cancel the payment, we will start returning the money to the payer’s account.
If the payment was made from a bank card, a YooMoney wallet, or via SberPay, the money will be refunded instantly.
If the payment was made using other payment methods, the process can take up to several days.

The response to the request will contain the `Payment` object with the current status.
```go
import (
    "github.com/rvinnie/yookassa-sdk-go/yookassa"   
    "github.com/rvinnie/yookassa-sdk-go/yookassa/common"
    "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
)

func main() {
    // Create a yookassa client by specifying the store ID and secret key
    yooclient := yookassa.NewClient('<Store ID>', '<Secret key>')
    // Create a payment handler
    paymentHandler := yookassa.NewPaymentHandler(yooclient)
    // Create a payment
    payment, _ := paymentHandler.CreatePayment(&yoopayment.Payment{
        Amount: &yoocommon.Amount{
            Value:    "1000.00",
            Currency: "RUB",
        },
        PaymentMethod: yoopayment.PaymentMethodType("bank_card"),
        Confirmation: yoopayment.Redirect{
            Type:      "redirect",
            ReturnURL: "https://www.example.com",
        },
        Description: "Test payment",
    })
    
    // We are waiting for the payment to be made 30 seconds
    time.Sleep(time.Second * 30)
    
    // Cancel payment
    payment, _ = paymentHandler.CancelPayment(payment.ID)
}
```
[Learn more about confirming and canceling payments](https://yookassa.ru/developers/payments/payment-process?lang=en#capture-and-cancel)

---

### Get payment information

[Information about the payment in the documentation](https://yookassa.ru/developers/api?lang=en#get_payment)

This request allows you to get the information about the current payment status by its unique ID.

The response to the request will contain the `Payment` object with the current status.

```go
package main

import "github.com/rvinnie/yookassa-sdk-go/yookassa"

func main() {
	// Create a yookassa client by specifying the store ID and secret key
	yooclient := yookassa.NewClient('<Store ID>', '<Secret key>')
	// Create a payment handler
	paymentHandler := yookassa.NewPaymentHandler(yooclient)
	// Getting payment information
	p, _ := paymentHandler.FindPayment("21b23b5b-000f-5061-a000-0674e49a8c10")
}
```
---

### Get a list of payments with filtering

[List of payments in documentation](https://yookassa.ru/developers/api?lang=en#get_payments_list)

The request allows you to receive the list of payments filtered by specified criteria.

The response will contain the list of payments with applied parameters of the request.
The list includes the information about the payments created within the last 3 years.
The list will be sorted by payment creation time in descending order.
If the number of results exceeds the value of limit, the list will be provided in fragments.
In this case, the response to the request will include a fragment of the list and the `next_cursor` parameter
with a `cursor` to the next fragment.

```go
package main

import (
	"github.com/rvinnie/yookassa-sdk-go/yookassa"
	"github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
)

func main() {
	// Create a yookassa client by specifying the store ID and secret key
	yooclient := yookassa.NewClient('<Store ID>', '<Secret key>')
	// Create a payment handler
	paymentHandler := yookassa.NewPaymentHandler(yooclient)
	
	//  Get all 'succeeded' payments of 3 per request
	var cursor string
	for {
		paymentsBatch, _ := paymentHandler.FindPayments(&yoopayment.PaymentListFilter{
			Limit:  3,
			Cursor: cursor,
			Status: yoopayment.Succeeded,
		})
		
		cursor = paymentsBatch.NextCursor
		// If the current piece of payments is the last one, exit the loop
		if cursor == "" { 
			break
		}
	}
}
```
[Learn more about working with lists](https://yookassa.ru/developers/using-api/lists?lang=en)