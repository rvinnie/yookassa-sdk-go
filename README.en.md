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
`go get github.com/rvinnie/yookassa-sdk-go`

### Commencing work
1. Import module
```golang
import "github.com/rvinnie/yookassa-sdk-go"
```
2. Configure a Client
```golang
import "github.com/rvinnie/yookassa-sdk-go"

func main() {
    client := yookassa.NewClient('<Account Id>', '<Secret Key>')	
}
```
3. Call the required API method. [More details in our documentation for the YooKassa API](https://yookassa.ru/developers/api?lang=en)

## Examples of using the API SDK
#### [YooKassa SDK Settings](https://github.com/rvinnie/yookassa-sdk-go/blob/main/docs/examples/01-configuration.md)
* [Authentication](https://github.com/rvinnie/yookassa-sdk-go/blob/main/docs/examples/01-configuration.en.md#Authentication)
* [Getting information about the store](https://github.com/rvinnie/yookassa-sdk-go/blob/main/docs/examples/01-configuration.en.md#Getting-information-about-the-store)
#### [Working with payments](https://github.com/rvinnie/yookassa-sdk-go/blob/main/docs/examples/01-payments.md)
* [Request to create a payment](https://github.com/rvinnie/yookassa-sdk-go/blob/main/docs/examples/02-payments.en.md#Request-to-create-a-payment)
* [Request for payment confirmation](https://github.com/rvinnie/yookassa-sdk-go/blob/main/docs/examples/02-payments.en.md#Request-for-payment-confirmation)
* [Request to cancel an incomplete payment](https://github.com/rvinnie/yookassa-sdk-go/blob/main/docs/examples/02-payments.en.md#Request-to-cancel-an-incomplete-payment)
* [Get payment information](https://github.com/rvinnie/yookassa-sdk-go/blob/main/docs/examples/02-payments.en.md#Get-payment-information)
* [Get a list of payments with filtering](https://github.com/rvinnie/yookassa-sdk-go/blob/main/docs/examples/02-payments.en.md#Get-a-list-of-payments-with-filtering)
