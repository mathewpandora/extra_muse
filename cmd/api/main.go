package main

import (
	"context"
	"extra_muse/internal/config"
	"extra_muse/internal/repository"
	"extra_muse/pkg/polza"
	"fmt"
	"log"
	"time"
)

func main() {

	configData := config.Load()
	dbPool, err := config.ConnectToDB(configData.DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	
	defer dbPool.Close() //при закрытиии руками наприменр это позволит всем транзакциям выполниться в бд 

 
	if err := repository.CreateGenerationsTablle(context.Background(),dbPool); err != nil {
		log.Fatal(err)
	}

	polzaClient := polza.New(configData.BASE_URL, configData.POLZA_API_KEY, dbPool)

	
	GenerateRequestSimple := polza.GenerateRequest{
		Model: "suno/generate",
		Input: polza.InputRequest{
			Prompt: "Реп про разработку на golang",
		},
	}


	res, err := polzaClient.Generate(GenerateRequestSimple)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.Message)

	time.Sleep(1000 * time.Second)
	
	fmt.Println("Отработали")
}
}