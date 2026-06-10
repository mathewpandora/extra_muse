package polza

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PolzaClient struct {
	//Клиент для пользы
	//все поля создаются один раз и больше не меняются
	DB *pgxpool.Pool
	BaseURL     string
	APIKey      string
	ContentType string
	HTTPClient  *http.Client //как пульт от телевизора (должен быть один экземпляр) Общий таймаут на весь запрос  Политика обработки редиректов Хранение кук
}

func New(BaseURL, APIKey string, dbPool *pgxpool.Pool) *PolzaClient {
	//валидация
	var ContentType = "application/json"
	return &PolzaClient{
		DB: dbPool,
		BaseURL: BaseURL,
		APIKey:  APIKey,
		HTTPClient: &http.Client{ //почитать
			//  ctx
			Timeout: 5 * time.Minute,
		},
		ContentType: ContentType,
	}
}



func (c *PolzaClient) sendRequest(method, url, token string, body io.Reader) (*http.Response, error) { //body - срез байтов в котором закодирован тело json ом
	
	Request, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, fmt.Errorf("sendRequest %w", err)
	}

	TokenString := fmt.Sprintf("Bearer %s", token)
	Request.Header.Set("Authorization", TokenString)
	Request.Header.Set("Content-Type", c.ContentType)

	resp, err := c.HTTPClient.Do(Request) //resp - указатель
	//wait group - группа ожиданий горутин

	if err != nil {
		return nil, fmt.Errorf("sendRequest: %w", err)
	}

	return resp, nil

}

// error группы
func (c *PolzaClient) Generate(r GenerateRequest) (*ShortResponse, error) {

	if r.Input.Prompt == "" {
		return nil, errors.New("Prompt cant be nill")
	}

	JsonBytes, err := json.Marshal(r)

	if err != nil {
		return nil, fmt.Errorf("Generate: %w", err)
	}

	Reader := bytes.NewReader(JsonBytes)

	go func() {
		fmt.Println("ASYNC START")
		resp, err := c.sendRequest("POST", c.BaseURL, c.APIKey, Reader)
		fmt.Println("ASYNC END")

	if err != nil {
		fmt.Println("Ошибка", err)
		return 
	}

	defer resp.Body.Close() //если не закрыть соединение не вернется в пул и со временем упадет
	
	BytesResponse, err := io.ReadAll(resp.Body) //он ожидает пакеты

	if err != nil {
		fmt.Println("Ошибка", err)
		return 
	}

	switch resp.StatusCode {
	case 200:
		var result GenerateResponse
		if err := json.Unmarshal(BytesResponse, &result); err != nil {
			fmt.Println("Generate, ", err)
		}
		if result.Status == "failed" {
			fmt.Println(errors.New(result.ErrorDetail.Message))
			return 
		}
		fmt.Printf("Response data: %+v\n", result) //+v возвращает структуру в человекочитаемом виде 
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		query := `
			INSERT INTO generations 
				(tg_id, prompt, status, gen_id, track_url1, image_url1, title1, 
				track_url2, image_url2, title2, error)
			VALUES 
				($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);
				`
	if	_, err := c.DB.Exec(ctx, query, nil, r.Input.Prompt, result.Status, result.ID, result.ResponseData[0].URL, result.ResponseData[0].Image, result.ResponseData[0].Title,
    result.ResponseData[0].URL, result.ResponseData[0].Image, result.ResponseData[0].Title, nil); err != nil {
		fmt.Println("ВАЖНО ЛОГИРОВАТЬ СИТУАЦИЯ КОГДА ГОРУТИНА НЕ ЗАПИСАЛА В БАЗУ")
	}
	fmt.Println(result)

	case 401:
		fmt.Println(errors.New("Ошибка авторизации. Проверьте API-ключ."))

	case 500:
		fmt.Println(errors.New("Сервис временно недоступен. Попробуйте позже."))

	default:
		fmt.Println(errors.New("other response error"))
	}
	}()

	return &ShortResponse{Message: "Запрос принят, ожидайте"}, nil
}




