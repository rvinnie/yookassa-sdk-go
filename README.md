[![Golang](https://img.shields.io/badge/Go-v1.19-EEEEEE?logo=go&logoColor=white&labelColor=00ADD8)](https://go.dev/)
[![License](https://img.shields.io/pypi/l/yookassa.svg)](LICENSE)

<div align="center">
    <h1 align="center">YooKassa API Golang Client Library
    </h1>
    <h3 align="center">Клиент для работы с платежами по <a href="https://yookassa.ru/developers/api">API ЮKassa</a>
    </h3>
    <p align="center">
        Russian | <a href="README.en.md">English</a> 
    </p>
</div>

### Установка
`go get github.com/rvinnie/yookassa-sdk-go`

### Начало работы
1. Импортируйте модуль
```golang
import "github.com/rvinnie/yookassa-sdk-go"
```
2. Установите данные для конфигурации
```golang
import "github.com/rvinnie/yookassa-sdk-go"

func main() {
    client := yookassa.NewClient('<Идентификатор магазина>', '<Секретный ключ>')	
}
```
3. Вызовите нужный метод API. [Подробнее в документации к API ЮKassa](https://yookassa.ru/developers/api)

## Примеры использования SDK
#### [Настройки SDK API ЮKassa](https://github.com/rvinnie/yookassa-sdk-go/blob/main/docs/examples/01-configuration.md)
* [Аутентификация](https://github.com/rvinnie/yookassa-sdk-go/blob/main/docs/examples/01-configuration.md#Аутентификация)
#### [Работа с платежами](https://github.com/rvinnie/yookassa-sdk-go/blob/main/docs/examples/01-payments.md)
* [Запрос на создание платежа](https://github.com/rvinnie/yookassa-sdk-go/blob/main/docs/examples/02-payments.md#Запрос-на-создание-платежа)
* [Запрос на подтверждение платежа](https://github.com/rvinnie/yookassa-sdk-go/blob/main/docs/examples/02-payments.md#Запрос-на-подтверждение-платежа)
* [Запрос на отмену незавершенного платежа](https://github.com/rvinnie/yookassa-sdk-go/blob/main/docs/examples/02-payments.md#Запрос-на-отмену-незавершенного-платежа)
* [Получить информацию о платеже](https://github.com/rvinnie/yookassa-sdk-go/blob/main/docs/examples/02-payments.md#Получить-информацию-о-платеже)
* [Получить список платежей с фильтрацией](https://github.com/rvinnie/yookassa-sdk-go/blob/main/docs/examples/02-payments.md#Получить-список-платежей-с-фильтрацией)