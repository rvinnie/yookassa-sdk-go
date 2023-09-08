## Работа с платежами

SDK позволяет создавать, подтверждать, отменять платежи, а также получать информацию о них.

* [Запрос на создание платежа](#Запрос-на-создание-платежа)
* [Запрос на подтверждение платежа](#Запрос-на-подтверждение-платежа)
* [Запрос на отмену незавершенного платежа](#Запрос-на-отмену-незавершенного-платежа)
* [Получить информацию о платеже](#Получить-информацию-о-платеже)
* [Получить список платежей с фильтрацией](#Получить-список-платежей-с-фильтрацией)

---

### Запрос на создание платежа
[Создание платежа в документации](https://yookassa.ru/developers/api?lang=bash#create_payment)

Чтобы принять оплату, необходимо создать объект платежа — `Payment`. Он содержит всю необходимую информацию
для проведения оплаты (сумму, валюту и статус). У платежа линейный жизненный цикл,
он последовательно переходит из статуса в статус.

В ответ на запрос придет объект платежа - `Payment` в актуальном статусе.

```go
import (
    "github.com/rvinnie/yookassa-sdk-go/yookassa"
    "github.com/rvinnie/yookassa-sdk-go/yookassa/common"
    "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
)

func main() {
	// Создаем yookassa клиента, указав идентификатор магазина и секретный ключ
	yooclient := yookassa.NewClient('<Идентификатор магазина>', '<Секретный ключ>')
	// Создаем обработчик платежей
	paymentHandler := yookassa.NewPaymentHandler(yooclient)
	// Создаем платеж
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

### Запрос на подтверждение платежа

[Подтверждение платежа в документации](https://yookassa.ru/developers/api?lang=bash#capture_payment)

Подтверждает вашу готовность принять платеж. После подтверждения платеж перейдет в статус succeeded.
Это значит, что вы можете выдать товар или оказать услугу пользователю.

Подтвердить можно только платеж в статусе `waiting_for_capture` и только в течение определенного времени
(зависит от способа оплаты). Если вы не подтвердите платеж в отведенное время, он автоматически перейдет
в статус `canceled`, и деньги вернутся пользователю.

В ответ на запрос придет объект платежа в актуальном статусе.

```go
import (
    "github.com/rvinnie/yookassa-sdk-go/yookassa"
    "github.com/rvinnie/yookassa-sdk-go/yookassa/common"
    "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
)

func main() {
	// Создаем yookassa клиента, указав идентификатор магазина и секретный ключ
	yooclient := yookassa.NewClient('<Идентификатор магазина>', '<Секретный ключ>')
	// Создаем обработчик платежей
	paymentHandler := yookassa.NewPaymentHandler(yooclient)
	// Создаем платеж
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

	// Ожидаем совершение платежа 30 секунд 
	time.Sleep(time.Second * 30)
	
	// Подтверждаем платеж
	payment, _ = paymentHandler.CapturePayment(payment)
}
```
[Подробнее о подтверждении и отмене платежей](https://yookassa.ru/developers/payments/payment-process#capture-and-cancel)


---

### Запрос на отмену незавершенного платежа
[Отмена платежа в документации](https://yookassa.ru/developers/api?lang=python#cancel_payment)

Отменяет платеж, находящийся в статусе `waiting_for_capture`. Отмена платежа значит, что вы не готовы
выдать пользователю товар или оказать услугу. Как только вы отменяете платеж, мы начинаем возвращать деньги на счет плательщика. Для платежей банковскими картами или из кошелька ЮMoney отмена происходит мгновенно. Для остальных способов оплаты возврат может занимать до нескольких дней.

В ответ на запрос придет объект платежа в актуальном статусе.
```go
import (
    "github.com/rvinnie/yookassa-sdk-go/yookassa"
    "github.com/rvinnie/yookassa-sdk-go/yookassa/common"
    "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
)

func main() {
    // Создаем yookassa клиента, указав идентификатор магазина и секретный ключ
    yooclient := yookassa.NewClient('<Идентификатор магазина>', '<Секретный ключ>')
    // Создаем обработчик платежей
    paymentHandler := yookassa.NewPaymentHandler(yooclient)
    // Создаем платеж
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
    
    // Ожидаем совершение платежа 30 секунд 
    time.Sleep(time.Second * 30)
    
    // Отменяем платеж
    payment, _ = paymentHandler.CancelPayment(payment.ID)
}
```
[Подробнее о подтверждении и отмене платежей](https://yookassa.ru/developers/payments/payment-process#capture-and-cancel)

---

### Получить информацию о платеже

[Информация о платеже в документации](https://yookassa.ru/developers/api?lang=bash#get_payment)

Запрос позволяет получить информацию о текущем состоянии платежа по его уникальному идентификатору.

В ответ на запрос придет объект платежа в актуальном статусе.

```go
package main

import "github.com/rvinnie/yookassa-sdk-go/yookassa"

func main() {
	// Создаем yookassa клиента, указав идентификатор магазина и секретный ключ
	yooclient := yookassa.NewClient('<Идентификатор магазина>', '<Секретный ключ>')
	// Создаем обработчик платежей
	paymentHandler := yookassa.NewPaymentHandler(yooclient)
	// Получаем информацию о платеже
	p, _ := paymentHandler.FindPayment("21b23b5b-000f-5061-a000-0674e49a8c10")
}
```
---

### Получить список платежей с фильтрацией

[Список платежей в документации](https://yookassa.ru/developers/api?lang=bash#get_payments_list)

Запрос позволяет получить список платежей, отфильтрованный по заданным критериям.

В ответ на запрос вернется список платежей с учетом переданных параметров. В списке будет информация о платежах,
созданных за последние 3 года. Список будет отсортирован по времени создания платежей в порядке убывания.

Если результатов больше, чем задано в `limit`, список будет выводиться фрагментами. В этом случае в ответе на запрос
вернется фрагмент списка и параметр `next_cursor` с указателем на следующий фрагмент.

```go
package main

import (
	"github.com/rvinnie/yookassa-sdk-go/yookassa"
	"github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
)

func main() {
	// Создаем yookassa клиента, указав идентификатор магазина и секретный ключ
	yooclient := yookassa.NewClient('<Идентификатор магазина>', '<Секретный ключ>')
	// Создаем обработчик платежей
	paymentHandler := yookassa.NewPaymentHandler(yooclient)
	
	//  Получаем все 'succeeded' платежи по 3 штуки за запрос
	var cursor string
	for {
		paymentsBatch, _ := paymentHandler.FindPayments(&yoopayment.PaymentListFilter{
			Limit:  3,
			Cursor: cursor,
			Status: yoopayment.Succeeded,
		})
		
		cursor = paymentsBatch.NextCursor
		// Если текущий кусок платежей последний, выходим из цикла
		if cursor == "" { 
			break
		}
	}
}
```
[Подробнее о работе со списками](https://yookassa.ru/developers/using-api/lists)