package main

import "net/http"

type PolzaClient struct {
	//Клиент для пользы
	//все поля создаются один раз и больше не меняются
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client //как пульт от телевизора (должен быть один экземпляр)
}

func New(BaseURL, APIKey string) *PolzaClient {
	//валидация
	return &PolzaClient{
		BaseURL:    BaseURL,
		APIKey:     APIKey,
		HTTPClient: &http.Client{},
	}
}

func (c *PolzaClient) Generate(r GenerateRequest)

type GenerateRequest struct {
	Style  string
	Title  string
	Prompt string
	Mode   string
}

func main() {

}
