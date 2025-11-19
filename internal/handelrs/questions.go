package handelrs

import (
	"QA-service/internal/models"
	"QA-service/internal/storage"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type QuestionHandler struct {
	storage *storage.Storage
}

func NewQuestionStorage(storage *storage.Storage) *QuestionHandler {
	return &QuestionHandler{storage: storage}
}

// Хэндлер для возвращения всех вопросов
func (qh *QuestionHandler) GetQuestions(w http.ResponseWriter, r *http.Request) {

	questions, err := qh.storage.GetQuestions()

	if err != nil {
		http.Error(w, "failed to get questions in handler", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)

}

// Хэдлер для создания нового вопроса
func (qh *QuestionHandler) CreateQuestion(w http.ResponseWriter, r *http.Request) {

	var question models.Question

	err := json.NewDecoder(r.Body).Decode(&question)
	if err != nil {
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		return
	}

	if question.Text == "" {
		log.Println("error: empty question text")
		http.Error(w, "can't create an empty question", http.StatusBadRequest)
		return
	}

	createdQuestion, err := qh.storage.CreateQuestion(question.Text)

	if err != nil {
		log.Println(fmt.Errorf("failed to create question in storage: %w", err))
		http.Error(w, "failed to create question", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdQuestion)

}

// Хэндлер получения вопроса по ID
func (qh *QuestionHandler) GetQuestion(w http.ResponseWriter, r *http.Request) {

	questionIDStr := chi.URLParam(r, "questionID")
	questionID, err := strconv.Atoi(questionIDStr)
	if err != nil {
		log.Println(fmt.Errorf("failed to get id from URL: %w", err))
		http.Error(w, "invalid question ID", http.StatusBadRequest)
		return
	}

	question, err := qh.storage.GetQuestion(questionID)
	if err != nil {
		log.Println(fmt.Errorf("failed to get question with ID: %w", err))
		http.Error(w, "No question with this ID", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(question)

}

func (qh *QuestionHandler) DeleteQuestion(w http.ResponseWriter, r *http.Request) {

	questionIDStr := chi.URLParam(r, "questionID")
	questionID, err := strconv.Atoi(questionIDStr)

	if err != nil {
		log.Println(fmt.Errorf("failed to get id from URL: %w", err))
		http.Error(w, "invalid question ID", http.StatusBadRequest)
		return
	}

	err = qh.storage.DeliteQuestion(questionID)
	if err != nil {
		log.Println(fmt.Errorf("failed to delite question: %w", err))
		http.Error(w, "failed to delite question", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)

}
