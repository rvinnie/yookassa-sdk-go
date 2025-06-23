## Работа с возвратами

С помощью SDK можно возвращать платежи — полностью или частично. Порядок возврата зависит от способа оплаты
(`payment_method`) исходного платежа. При оплате банковской картой деньги возвращаются на карту,
которая была использована для проведения платежа. [Как проводить возвраты](https://yookassa.ru/developers/payments/refunds?lang=ru)

Часть способов оплаты (например, наличные) не поддерживают возвраты. [Какие платежи можно вернуть](https://yookassa.ru/developers/payment-methods/overview#all?lang=ru)

* [Запрос на создание возврата](#Запрос-на-создание-возврата)
* [Получить информацию о возврате](#Получить-информацию-о-возврате)
* [Получить список возвратов с фильтрацией](#Получить-список-возвратов-с-фильтрацией)

---

### Запрос на создание возврата

[Создание возврата в документации](https://yookassa.ru/developers/api?lang=bash#create_refund?lang=ru)  

Создает возврат успешного платежа на указанную сумму. Платеж можно вернуть только в течение трех лет с момента его создания.
Комиссия ЮKassa за проведение платежа не возвращается.

В ответ на запрос придет объект `Refund` в актуальном статусе.

```go
import (
    "github.com/imgrigorev/yookassa-sdk-go/yookassa"
    "github.com/imgrigorev/yookassa-sdk-go/yookassa/refund"
)

func main() {
    // Создаем yookassa клиента, указав идентификатор магазина и секретный ключ
    yooclient := yookassa.NewClient('<Store ID>', '<Secret key>')
    // Создаем обработчик возвратов
    refundHandler := yookassa.NewRefundHandler(client)
    // Создаем возврат
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

### Получить информацию о возврате

[Информация о возврате в документации](https://yookassa.ru/developers/api?lang=bash#get_refund?lang=ru)

Запрос позволяет получить информацию о текущем состоянии возврата по его уникальному идентификатору.

В ответ на запрос придет объект `Refund` в актуальном статусе.

```go
import (
    "github.com/imgrigorev/yookassa-sdk-go/yookassa"
    "github.com/imgrigorev/yookassa-sdk-go/yookassa/refund"
)

func main() {
    // Создаем yookassa клиента, указав идентификатор магазина и секретный ключ
    yooclient := yookassa.NewClient('<Store ID>', '<Secret key>')
    // Создаем обработчик возвратов 
    refundHandler := yookassa.NewRefundHandler(client)
    // Получаем объект возврата
    refund, _ := refundHandler.FindRefund("2c87b72c-0015-5000-9000-172b6038152a")
}
```

---

### Получить список возвратов с фильтрацией

[Список возвратов в документации](https://yookassa.ru/developers/api?lang=bash#get_refunds_list?lang=ru)

Запрос позволяет получить список возвратов, отфильтрованный по заданным критериям.

В ответ на запрос вернется список возвратов с учетом переданных параметров. В списке будет информация о возвратах,
созданных за последние 3 года. Список будет отсортирован по времени создания возвратов в порядке убывания.

Если результатов больше, чем задано в `limit`, список будет выводиться фрагментами. В этом случае в ответе на запрос
вернется фрагмент списка и параметр `next_cursor` с указателем на следующий фрагмент.

```go
import (
    "github.com/imgrigorev/yookassa-sdk-go/yookassa"
    "github.com/imgrigorev/yookassa-sdk-go/yookassa/refund"
)

func main() {
    // Создаем yookassa клиента, указав идентификатор магазина и секретный ключ
    yooclient := yookassa.NewClient('<Store ID>', '<Secret key>')
    // Создаем обработчик возвратов 
    refundHandler := yookassa.NewRefundHandler(client)
    // Получаем список объектов возврата (последние 3 со статусом succeeded)
    refunds, err := refundHandler.FindRefunds(&yoorefund.RefundListFilter{
        Status: yoorefund.Succeeded,
        Limit:  3,
    })
}
```