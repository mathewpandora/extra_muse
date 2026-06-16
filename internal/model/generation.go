package model

type Track struct {
	URL      string
	Image    string
	Title    string
	Duration float32
}

type NewGenerationData struct {
	TgID      int64
	Prompt    string `json:"prompt"`
	Style     string `json:"style"`
	Title     string `json:"title"`
	GenID     string `json:"gen_id"`
	TrackURL1 string `json:"track_url1"`
	ImageURL1 string `json:"image_url1"`
	Title1    string `json:"title1"`
	TrackURL2 string `json:"track_url2"`
	ImageURL2 string `json:"image_url2"`
	Title2    string `json:"title2"`
	Status    string `json:"status"`
	Error     string `json:"error"`
}
