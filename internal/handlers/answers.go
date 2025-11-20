package handlers

import (
	"QA-service/internal/models"
	"QA-service/internal/storage"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type AnswerHandler struct {
	storage *storage.Storage
}

func NewAnswerHandler(storage *storage.Storage) *AnswerHandler {
	return &AnswerHandler{storage: storage}
}

// Хэндлер для создания ответа на вопрос с указанным ID
func (ah *AnswerHandler) CreateAnswer(w http.ResponseWriter, r *http.Request) {

	questionIDStr := chi.URLParam(r, "questionID")
	questionID, err := strconv.Atoi(questionIDStr)

	if err != nil {
		log.Println(fmt.Errorf("failed to get question ID from URL: %w", err))
		http.Error(w, "failed to find question with this ID", http.StatusBadRequest)
		return
	}

	var answer models.Answer

	err = json.NewDecoder(r.Body).Decode(&answer)

	if err != nil {
		log.Println(fmt.Errorf("failed to decode json from answer body: %w", err))
		http.Error(w, "invalid answer body", http.StatusBadRequest)
		return
	}

	if answer.Text == "" {
		http.Error(w, "Text required", http.StatusBadRequest)
		return
	}

	if answer.UserID == "" {
		http.Error(w, "ID required", http.StatusBadRequest)
		return
	}

	createdAnswer, err := ah.storage.CreateAnswer(questionID, answer.UserID, answer.Text)

	if err != nil {
		log.Println(fmt.Errorf("failed to create answer: %w", err))
		http.Error(w, "failed to create answer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdAnswer)

	log.Printf("New answer created; question ID = %d created answer ID = %d, text = %s", questionID, createdAnswer.ID, createdAnswer.Text)
}

// Хэндлер для возвращения конкретного ответа по указаному ID ответа
func (ah *AnswerHandler) GetAnswer(w http.ResponseWriter, r *http.Request) {
	answerIDStr := chi.URLParam(r, "answerID")
	answerID, err := strconv.Atoi(answerIDStr)

	if err != nil {
		log.Println(fmt.Errorf("failed to get id from URL: %w", err))
		http.Error(w, "invalid answer ID", http.StatusBadRequest)
		return
	}

	answer, err := ah.storage.GetAnswer(answerID)

	if err != nil {
		log.Println(fmt.Errorf("failed to get answer with ID: %w", err))
		http.Error(w, "No answer with this ID", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(answer)

	log.Printf("user got answer = %d, text = %s", answer.ID, answer.Text)

}

// Хэндлер для удаления ответа с указанным ID
func (ah *AnswerHandler) DeleteAnswer(w http.ResponseWriter, r *http.Request) {

	answerIDStr := chi.URLParam(r, "answerID")
	answerID, err := strconv.Atoi(answerIDStr)

	if err != nil {
		log.Println(fmt.Errorf("failed to get id from URL: %w", err))
		http.Error(w, "invalid answer ID", http.StatusBadRequest)
		return
	}

	err = ah.storage.DeleteAnswer(answerID)

	if err != nil {
		log.Println(fmt.Errorf("failed to delite answer: %w", err))
		http.Error(w, "failed to delite answer", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)

	log.Printf("Deleted answer with ID = %d", answerID)
}
