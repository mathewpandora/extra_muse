package worker

//Producer-Consumer — классическое название из учебников
//Worker Pool — то же самое, но с акцентом на пул воркеров
import (
	"extra_muse/internal/model"
	"extra_muse/internal/repository"
	"extra_muse/pkg/polza"
	"log"
)

//Что делает воркер:

// Принимает данные для запроса
//Отправляет запрос и ждет
//Принимает и кладет в бд

//Воркер пул

//буферизированный на 100 канал с инпут данными с ручки

type Pool struct {
	Jobs chan model.NewGenerationData
}

type Worker struct {
	Pool                 Pool
	GenerationRepository repository.GenerationRepository
	PolzaClient          polza.PolzaClient
}

func (w *Worker) Work() {

	for {
		GenData := <-w.Pool.Jobs

		resp, err := w.PolzaClient.Generate(

			polza.GenerateRequest{
				Model: "suno/generate",
				Input: polza.InputRequest{
					Prompt: GenData.Prompt,
					Style:  GenData.Style,
					Title:  GenData.Title,
				}})

		if err != nil {
			saveError := w.GenerationRepository.Save(
				model.NewGenerationData{
					TgID:   GenData.TgID,
					Prompt: GenData.Prompt,
					Style:  GenData.Style,
					Title:  GenData.Title,
					Status: "error",
					Error:  err.Error(),
				},
			)

			if saveError != nil {
				log.Printf("failed to save error to DB: %v", saveError)
			}
			continue
		}

		w.GenerationRepository.Save(
			model.NewGenerationData{
				TgID:      GenData.TgID,
				Prompt:    GenData.Prompt,
				Style:     GenData.Style,
				Title:     GenData.Title,
				Status:    "done",
				TrackURL1: resp.ResponseData[0].URL,
				ImageURL1: resp.ResponseData[0].Image,
				Title1:    resp.ResponseData[0].Title,
				TrackURL2: resp.ResponseData[1].URL,
				ImageURL2: resp.ResponseData[1].Image,
				Title2:    resp.ResponseData[1].Title,
			},
		)
	}

}

type WorkerPool struct {
	Pool         Pool
	WorkerAmount int64
}

func NewWorkerPool(n, m int64) *WorkerPool {
	return &WorkerPool{
		Pool: Pool{
			Jobs: make(chan model.NewGenerationData, m),
		},
		WorkerAmount: n,
	}
}

func (wp *WorkerPool) Start(gr repository.GenerationRepository, pc polza.PolzaClient) {

	for range wp.WorkerAmount {

		newWorker := Worker{
			Pool:                 wp.Pool,
			GenerationRepository: gr,
			PolzaClient:          pc, // ← запятая нужна
		}

		go newWorker.Work()

	}
}

func (wp *WorkerPool) Add(Job model.NewGenerationData) {
	wp.Pool.Jobs <- Job
}
