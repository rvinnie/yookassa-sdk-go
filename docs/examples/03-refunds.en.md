## Working with refunds

Using the SDK, you can return payments - in full or in part. The return procedure depends on the payment method.
(`payment_method`) of the original payment. When paying with a bank card, the money is returned to the card,
that was used to make the payment. [How to make refunds](https://yookassa.ru/developers/payment-acceptance/after-the-payment/refunds?lang=en)

Some payment methods (for example, cash) do not support refunds. [What payments can be refunded](https://yookassa.ru/developers/payment-methods/overview#all?lang=en)

* [Request to create a refund](#Request-to-create-a-refund)
* [Get refund information](#Get-refund-information)
* [Get a list of refunds with filtering](#Get-a-list-of-refunds-with-filtering)

---

### Request to create a refund

[Creating a refund in documentation](https://yookassa.ru/developers/api?lang=bash#create_refund?lang=en)

Generates a successful payment refund for the specified amount. A payment can only be returned within three years of its creation.
YooKassa commission for making a payment is non-refundable.

In response to the request, the `Refund` object will come in the current status.

```go
import (
    "github.com/rvinnie/yookassa-sdk-go/yookassa"
    "github.com/rvinnie/yookassa-sdk-go/yookassa/refund"
)

func main() {
    // Create a yookassa client by specifying the store ID and secret key
    yooclient := yookassa.NewClient('<Store ID>', '<Secret key>')
    // Create a refund handler
    refundHandler := yookassa.NewRefundHandler(client)
    // Create a refund
    refund, err := refundHandler.CreateRefund(&yoorefund.Refund{
        PaymentId: "2c79414f-000f-5100-9000-1d082dd142ea",
        Amount: &yoocommon.Amount{
            Value:    "123",
            Currency: "RUB",
        },
        Description: "Test refund :)",
    })
}
```

---

### Get refund information

[Information about the refund in the documentation](https://yookassa.ru/developers/api?lang=bash#get_refund?lang=en)

The request allows you to get information about the current state of the refund by its unique identifier.

In response to the request, the `Refund` object will come in the current status.

```go
import (
    "github.com/rvinnie/yookassa-sdk-go/yookassa"
    "github.com/rvinnie/yookassa-sdk-go/yookassa/refund"
)

func main() {
    // Create a yookassa client by specifying the store ID and secret key
    yooclient := yookassa.NewClient('<Store ID>', '<Secret key>')
    // Create a refund handler 
    refundHandler := yookassa.NewRefundHandler(client)
    // Get the refund object
    refund, _ := refundHandler.FindRefund("2c87b72c-0015-5000-9000-172b6038152a")
}
```

---

### Get a list of refunds with filtering

[list of refunds in the documentation](https://yookassa.ru/developers/api?lang=bash#get_refunds_list?lang=en)

The request allows you to get a list of refunds filtered by the specified criteria.

In response to the request, a list of refunds will be returned, taking into account the passed parameters. The list will contain information about refunds,
created in the last 3 years. The list will be sorted by refund creation time in descending order.

If there are more results than specified in `limit`, the list will be displayed in fragments. In this case, in response to a request
the list fragment and the `next_cursor` parameter with a pointer to the next fragment will be returned.

```go
import (
    "github.com/rvinnie/yookassa-sdk-go/yookassa"
    "github.com/rvinnie/yookassa-sdk-go/yookassa/refund"
)

func main() {
    // Create a yookassa client by specifying the store ID and secret key
    yooclient := yookassa.NewClient('<Store ID>', '<Secret key>')
    // Create a refund handler 
    refundHandler := yookassa.NewRefundHandler(client)
    // We get a list of payment objects (the last 3 with the status succeeded)
    refunds, err := refundHandler.FindRefunds(&yoorefund.RefundListFilter{
        Status: yoorefund.Succeeded,
        Limit:  3,
    })
}
```