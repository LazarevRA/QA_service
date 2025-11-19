package storage

import (
	models "QA-service/internal/models"

	"time"
)

// PostQuestion создает новый вопрос
func (s *Storage) CreateQuestion(text string) (*models.Question, error) {
	question := &models.Question{
		Text:      text,
		CreatedAt: time.Now(),
	}

	err := s.DB.Create(question).Error
	if err != nil {
		return nil, err
	}

	return question, nil
}

// GetQuestions возвращает все вопросы
func (s *Storage) GetQuestions() ([]models.Question, error) {
	var questions []models.Question
	err := s.DB.Find(&questions).Error
	return questions, err
}

// GetQuestion возвращает вопрос по ID и все ответы на него
func (s *Storage) GetQuestion(id uint) (*models.Question, error) {
	var question models.Question
	err := s.DB.Preload("Answers").First(&question, id).Error
	if err != nil {
		return nil, err
	}
	return &question, nil
}

// DeliteQuestion удаляет вопрос и все ответы на него
func (s *Storage) DeliteQuestion(id uint) error {
	return s.DB.Delete(&models.Question{}, id).Error
}
