package main

import (
	"extra_muse/polza"
	"fmt"
	"os"
)

func main() {

	API_KEY := os.Getenv("POLZA_API_KEY")

	PolzaClient := polza.New("https://polza.ai/api/v1/media", API_KEY)

	GenerateRequestSimple := polza.GenerateRequest{
		Model: "suno/generate",
		Input: polza.InputRequest{
			Prompt: "Электронная музыка. Хардкор техно. Бочки. Басс. Быстро. ",
		},
	}

	res, err := PolzaClient.Generate(GenerateRequestSimple)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}
