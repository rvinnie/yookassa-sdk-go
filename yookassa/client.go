package yookassa

type Client struct {
	ApiUrl    string
	AccountId string
	SecretKey string
}

func NewClient(accountId string, secretKey string) *Client {
	return &Client{
		ApiUrl:    "https://api.yookassa.ru/v3/",
		AccountId: accountId,
		SecretKey: secretKey,
	}
}
