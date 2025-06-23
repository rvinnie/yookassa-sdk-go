[![Golang](https://img.shields.io/badge/Go-v1.19-EEEEEE?logo=go&logoColor=white&labelColor=00ADD8)](https://go.dev/)
[![License](https://img.shields.io/pypi/l/yookassa.svg)](LICENSE)

<div align="center">
    <h1 align="center">YooKassa API Golang Client Library
    </h1>
    <h3 align="center">This product is used for managing payments under <a href="https://yookassa.ru/developers/api?lang=en">The YooKassa API</a>
    </h3>
    <p align="center">
        <a href="README.md">Russian</a> | English 
    </p>
</div>

### Installation
`go get github.com/imgrigorev/yookassa-sdk-go`

### Commencing work
1. Import module
```golang
import "github.com/imgrigorev/yookassa-sdk-go"
```
2. Configure a Client
```golang
import "github.com/imgrigorev/yookassa-sdk-go"

func main() {
    client := yookassa.NewClient('<Account Id>', '<Secret Key>')	
}
```
3. Call the required API method. [More details in our documentation for the YooKassa API](https://yookassa.ru/developers/api?lang=en)

## Examples of using the API SDK
#### [YooKassa SDK Settings](https://github.com/imgrigorev/yookassa-sdk-go/blob/main/docs/examples/01-configuration.en.md)
* [Authentication](https://github.com/imgrigorev/yookassa-sdk-go/blob/main/docs/examples/01-configuration.en.md#Authentication)
* [Getting information about the store](https://github.com/imgrigorev/yookassa-sdk-go/blob/main/docs/examples/01-configuration.en.md#Getting-information-about-the-store)
#### [Working with payments](https://github.com/imgrigorev/yookassa-sdk-go/blob/main/docs/examples/02-payments.en.md)
* [Request to create a payment](https://github.com/imgrigorev/yookassa-sdk-go/blob/main/docs/examples/02-payments.en.md#Request-to-create-a-payment)
* [Request for payment confirmation](https://github.com/imgrigorev/yookassa-sdk-go/blob/main/docs/examples/02-payments.en.md#Request-for-payment-confirmation)
* [Request to cancel an incomplete payment](https://github.com/imgrigorev/yookassa-sdk-go/blob/main/docs/examples/02-payments.en.md#Request-to-cancel-an-incomplete-payment)
* [Get payment information](https://github.com/imgrigorev/yookassa-sdk-go/blob/main/docs/examples/02-payments.en.md#Get-payment-information)
* [Get a list of payments with filtering](https://github.com/imgrigorev/yookassa-sdk-go/blob/main/docs/examples/02-payments.en.md#Get-a-list-of-payments-with-filtering)
#### [Working with refunds](https://github.com/imgrigorev/yookassa-sdk-go/blob/main/docs/examples/03-refunds.en.md)
* [Request to create a refund](https://github.com/imgrigorev/yookassa-sdk-go/blob/main/docs/examples/03-refunds.en.md#Request-to-create-a-refund)
* [Get refund information](https://github.com/imgrigorev/yookassa-sdk-go/blob/main/docs/examples/03-refunds.en.md#Get-refund-information)
* [Get a list of refunds with filtering](https://github.com/imgrigorev/yookassa-sdk-go/blob/main/docs/examples/03-refunds.en.md#Get-a-list-of-refunds-with-filtering)
#### [Working with Webhooks](https://github.com/imgrigorev/yookassa-sdk-go/blob/main/docs/examples/04-webhooks.en.md)
* [Example Webhook Processing](https://github.com/imgrigorev/yookassa-sdk-go/blob/main/docs/examples/04-webhooks.en.md#Example-Webhook-Processing)
* [Local Testing](https://github.com/imgrigorev/yookassa-sdk-go/blob/main/docs/examples/04-webhooks.en.md#Local-Testing)
