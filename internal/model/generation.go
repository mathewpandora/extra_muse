package model

type Track struct{
	URL string 
	Image string 
	Title string 
	Duration float32 
}

type NewGenerationData struct{
	TgID int64
	Prompt string `json:"prompt"`
	Style  string `json:"style"`
	Title  string `json:"title"`
	
}

