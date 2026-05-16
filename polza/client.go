package polza

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type PolzaClient struct {
	//Клиент для пользы
	//все поля создаются один раз и больше не меняются
	BaseURL     string
	APIKey      string
	ContentType string
	HTTPClient  *http.Client //как пульт от телевизора (должен быть один экземпляр) Общий таймаут на весь запрос  Политика обработки редиректов Хранение кук
}

func New(BaseURL, APIKey string) *PolzaClient {
	//валидация
	var ContentType = "application/json"
	return &PolzaClient{
		BaseURL: BaseURL,
		APIKey:  APIKey,
		HTTPClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
		ContentType: ContentType,
	}
}

func (c *PolzaClient) sendRequest(method, url, token string, body io.Reader) (*http.Response, error) { //body - срез байтов в котором закодирован тело json ом

	Request, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, errors.New("Request has not built")
	}

	TokenString := fmt.Sprintf("Bearer %s", token)
	Request.Header.Set("Authorization", TokenString)
	Request.Header.Set("Content-Type", c.ContentType)

	resp, err := c.HTTPClient.Do(Request) //resp - указатель

	if err != nil {
		return nil, fmt.Errorf("Ошибка: %w", err)
	}

	return resp, nil

}

func (c *PolzaClient) Generate(r GenerateRequest) (*GenerateResponse, error) {

	if r.Input.Prompt == "" {
		return nil, errors.New("Prompt cant be nill")

	}

	JsonBytes, err := json.Marshal(r)

	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}

	Reader := bytes.NewReader(JsonBytes)

	resp, err := c.sendRequest("POST", c.BaseURL, c.APIKey, Reader)

	if err != nil {
		return nil, fmt.Errorf("sendRequest: %w", err)
		//тут сетевые ошибки
	}

	//после завершения функции поток закроется и его нельзя будет читать
	defer resp.Body.Close() //если не закрыть соединение не вернется в пул и со временем упадет

	BytesResponse, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %w", err)
	}

	switch resp.StatusCode {

	case 200:

		var result GenerateResponse

		if err := json.Unmarshal(BytesResponse, &result); err != nil {
			return nil, fmt.Errorf("Generate Method, %w", err)
		}

		if result.Status == "failed" {
			return &result, errors.New(result.ErrorDetail.Message)
		}
		return &result, nil

	case 401:

		return nil, errors.New("Response 401")

	case 500:

		return nil, errors.New("Response 500")

	default:
		return nil, nil
	}

}
