## Настройки SDK API ЮKassa

[Справочник API ЮKassa](https://yookassa.ru/developers/api)
* [Аутентификация](#Аутентификация)
---

### Аутентификация

Для работы с API необходимо создать клиента, указав идентификатор магазина и секретный ключ.

```go
import "github.com/rvinnie/yookassa-sdk-go"

func main() {
    client := yookassa.NewClient('<Идентификатор магазина>', '<Секретный ключ>')	
}
```