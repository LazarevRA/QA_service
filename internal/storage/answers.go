package storage

import (
	"QA-service/internal/models"
	"fmt"
	"time"
)

// CreateAnswer создает ответ на вопрос
func (s *Storage) CreateAnswer(questionID uint, userID, text string) (*models.Answer, error) {
	var question models.Question

	//Проверка существования вопроса
	err := s.DB.First(&question, questionID).Error
	if err != nil {
		return nil, fmt.Errorf("question with this ID not found")
	}

	answer := &models.Answer{
		QuestionID: questionID,
		UserID:     userID,
		Text:       text,
		CreatedAt:  time.Now(),
	}

	err = s.DB.Create(answer).Error
	if err != nil {
		return nil, err
	}
	return answer, nil
}

// GetAnswer возвращает конкретный ответ по ID
func (s *Storage) GetAnswer(id uint) (*models.Answer, error) {
	var answer models.Answer
	err := s.DB.First(&answer, id).Error
	if err != nil {
		return nil, err
	}
	return &answer, nil
}

// DeleteAnswer удаляет ответ по ID
func (s *Storage) DeleteAnswer(id uint) error {
	return s.DB.Delete(&models.Answer{}, id).Error
}
