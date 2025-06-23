## YooKassa SDK Settings

[YooKassa API Reference](https://yookassa.ru/developers/api?lang=en)
* [Authentication](#Authentication)
* [Getting information about the store](#Getting-information-about-the-store)
---

### Authentication

To work with the API, you need to create a client, specifying the store ID and secret key.

```go
import "github.com/imgrigorev/yookassa-sdk-go"

func main() {
    client := yookassa.NewClient('<Store ID>', '<Secret key>')	
}
```

---

### Getting information about the store

After setting the configuration, you can check the correctness of the data, as well as get information about the store.

```go
import "github.com/imgrigorev/yookassa-sdk-go/yookassa"

func main() {
    // Create a yookassa client by specifying the store ID and secret key
    yooclient := yookassa.NewClient('<Store ID>', '<Secret key>')
    // Create a settings handler
	settingsHandler := yookassa.NewSettingsHandler(yooclient)
    // Get information about store or gateway settings
	settings, _ := settingsHandler.GetAccountSettings(nil)
}
```
As a result, we will see something like this:
```
#0 dict(5) 
    ['account_id'] => str(6) "XXXXXX"
    ['test'] => bool(True) 
    ['fiscalization_enabled'] => bool(True) 
    ['payment_methods'] => list(2) 
        [0] => str(9) "yoo_money"
        [1] => str(9) "bank_card"
    ['status'] => str(7) "enabled"
```
Read more about the settings object in the [API documentation](https://yookassa.ru/developers/api?lang=en#me_object)
