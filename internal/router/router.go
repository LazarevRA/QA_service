package router

import (
	"QA-service/internal/handlers"
	"QA-service/internal/storage"

	"github.com/go-chi/chi/v5"
)

func NewRouter(storage *storage.Storage) *chi.Mux {
	r := chi.NewRouter()

	quetionHandler := handlers.NewQuestionHandler(storage)
	answerHandler := handlers.NewAnswerHandler(storage)

	r.Route("/questions", func(r chi.Router) {
		r.Get("/", quetionHandler.GetQuestions)
		r.Post("/", quetionHandler.CreateQuestion)

		r.Route("/{questionID}", func(r chi.Router) {
			r.Get("/", quetionHandler.GetQuestion)
			r.Delete("/", quetionHandler.DeleteQuestion)
			r.Post("/answers/", answerHandler.CreateAnswer)
		})
	})

	r.Route("/answers", func(r chi.Router) {
		r.Get("/{answerID}", answerHandler.GetAnswer)
		r.Delete("/{answerID}", answerHandler.DeleteAnswer)
	})

	return r
}
