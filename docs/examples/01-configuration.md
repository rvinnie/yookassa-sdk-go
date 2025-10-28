## Настройки SDK API ЮKassa

[Справочник API ЮKassa](https://yookassa.ru/developers/api)
* [Аутентификация](#Аутентификация)
* [Получение информации о магазине](#Получение-информации-о-магазине)
---

### Аутентификация

Для работы с API необходимо создать клиента, указав идентификатор магазина и секретный ключ.

```go
import "github.com/rvinnie/yookassa-sdk-go/yookassa"

func main() {
    client := yookassa.NewClient('<Идентификатор магазина>', '<Секретный ключ>')	
}
```

---

### Получение информации о магазине

После установки конфигурации можно проверить корректность данных, а также получить информацию о магазине.

```go
import "github.com/rvinnie/yookassa-sdk-go/yookassa"

func main() {
    // Создаем yookassa клиента, указав идентификатор магазина и секретный ключ
    yooclient := yookassa.NewClient('<Идентификатор магазина>', '<Секретный ключ>')
    // Создаем обработчик настроек
    settingsHandler := yookassa.NewSettingsHandler(yooclient)
    // Получаем информацию о настройках магазина или шлюза
    settings, _ := settingsHandler.GetAccountSettings(nil)
}
```
В результате мы увидим примерно следующее:
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
Подробнее про объект настроек в [документации к API](https://yookassa.ru/developers/api#me_object)
