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
	ID          string      `json:"id"` //gen_2172375911287230465
	Status      string      `json:"status"`//completed
	Created     int         `json:"created"`//1780238852 (привести в нормальный формат)
	ResponseData []Track  `json:"data"`
	ErrorDetail ErrorDetail `json:"error"`
}

type Track struct{
	URL string `json:"url"`
	Image string `json:"thumbnail_url"`
	Title string `json:"title"`
	Duration float32  `json:"duration"`
}

type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ShortResponse struct {
	Message string
}
