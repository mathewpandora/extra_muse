package polza

type GenerateRequest struct {
	Model string       `json:"model"`
	Input InputRequest `json:"input"`
	//нужны ли тут конструкторы
}

type InputRequest struct {
	Prompt string `json:"prompt"`
	Style  string `json:"style"`
	Title  string `json:"title"`
}

type GenerateResponse struct {
	ID          string      `json:"id"`
	Object      string      `json:"object"`
	Status      string      `json:"status"`
	Created     int         `json:"created"`
	Model       string      `json:"model"`
	ErrorDetail ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
