package main

import (
	"QA-service/internal/config"
	"QA-service/internal/storage"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("Starting Q&A Service...")

	// Загружаем конфигурацию
	cfg := config.Load()

	// Инициализируем хранилище
	storage, err := storage.NewStorage(cfg.GetDSN())
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}
	defer storage.Close()

	// Проверяем подключение к БД
	if err := storage.HealthCheck(); err != nil {
		log.Fatalf("Database health check failed: %v", err)
	}

	log.Println("Successfully connected to database!")

	// ТЕСТ: Создаем тестовый вопрос
	question, err := storage.CreateQuestion("Как работает Go?")
	if err != nil {
		log.Printf("Failed to create test question: %v", err)
	} else {
		log.Printf("Test question created with ID: %d", question.ID)
	}

	// ТЕСТ: Получаем все вопросы
	questions, err := storage.GetQuestions()
	if err != nil {
		log.Printf("Failed to get questions: %v", err)
	} else {
		log.Printf("Found %d questions", len(questions))
		for _, q := range questions {
			log.Printf("Question %d: %s", q.ID, q.Text)
		}
	}

	// ТЕСТ: Создаем ответ
	if question != nil {
		answer, err := storage.CreateAnswer(
			question.ID,
			"550e8400-e29b-41d4-a716-446655440000",
			"Go работает очень эффективно!",
		)
		if err != nil {
			log.Printf("Failed to create test answer: %v", err)
		} else {
			log.Printf("Test answer created with ID: %d", answer.ID)
		}

		// ТЕСТ: Получаем вопрос с ответами
		questionWithAnswers, err := storage.GetQuestion(question.ID)
		if err != nil {
			log.Printf("Failed to get question with answers: %v", err)
		} else {
			log.Printf("Question %d has %d answers",
				questionWithAnswers.ID, len(questionWithAnswers.Answers))
			for _, a := range questionWithAnswers.Answers {
				log.Printf(" - Answer %d: %s (user: %s)", a.ID, a.Text, a.UserID)
			}
		}
	}

	log.Println("Service is ready!")
}
