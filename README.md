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