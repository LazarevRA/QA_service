package handlers

import (
	"QA-service/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type MockStorage struct {
	questions []models.Question
	nextID    int
}

func NewMockStorage() *MockStorage {
	return &MockStorage{
		questions: []models.Question{},
		nextID:    1,
	}
}

func (m *MockStorage) GetQuestions() ([]models.Question, error) {
	return m.questions, nil
}

func (m *MockStorage) CreateQuestion(text string) (*models.Question, error) {
	question := &models.Question{
		ID:   m.nextID,
		Text: text,
	}
	m.questions = append(m.questions, *question)
	m.nextID++
	return question, nil
}

func (m *MockStorage) GetQuestion(id int) (*models.Question, error) {
	for _, q := range m.questions {
		if q.ID == id {
			return &q, nil
		}
	}
	return nil, nil
}

func (m *MockStorage) DeleteQuestion(id int) error {
	for i, q := range m.questions {
		if q.ID == id {
			m.questions = append(m.questions[:i], m.questions[i+1:]...)
			return nil
		}
	}
	return nil
}

// Тест получения всех вопросов
func TestGetQuestions(t *testing.T) {
	mockStorage := NewMockStorage()
	mockStorage.CreateQuestion("Question 1")
	mockStorage.CreateQuestion("Question 2")

	handler := NewQuestionHandler(mockStorage)

	req, err := http.NewRequest("GET", "/questions", nil)
	require.NoError(t, err, "Failed to create request")

	rr := httptest.NewRecorder()
	handler.GetQuestions(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Wrong status code")

	var questions []models.Question

	err = json.Unmarshal(rr.Body.Bytes(), &questions)

	require.NoError(t, err, "Failed to unmarshal response JSON")

	assert.Len(t, questions, 2, "Wrong number of questions")

	assert.Equal(t, "Question 1", questions[0].Text, "First question text doesn't match")
	assert.Equal(t, "Question 2", questions[1].Text, "Second question text doesn't match")
}

// Тест создания вопроса с корректным телом
func TestCreateQuestion_OK(t *testing.T) {
	mockStorage := NewMockStorage()
	handler := NewQuestionHandler(mockStorage)

	questionText := map[string]string{"text": "Test question"}
	jsonText, err := json.Marshal(questionText)
	require.NoError(t, err, "Failed to marshal text")

	req, err := http.NewRequest("POST", "/questions", bytes.NewBuffer(jsonText))
	require.NoError(t, err, "Failed to create request")

	rr := httptest.NewRecorder()
	handler.CreateQuestion(rr, req)

	var question models.Question

	err = json.Unmarshal(rr.Body.Bytes(), &question)
	require.NoError(t, err, "Failed to unmarshal text")

	assert.Equal(t, "Test question", question.Text, "Question text doesn't match")
	assert.Equal(t, 1, question.ID, "Question ID should be 1")
}

// Тест создания вопроса с пустым телом
func TestCreateQuestion_Empty(t *testing.T) {
	mockStorage := NewMockStorage()
	handler := NewQuestionHandler(mockStorage)

	questionText := map[string]string{"text": ""}
	jsonText, err := json.Marshal(questionText)

	require.NoError(t, err, "Failed to marshal request body")

	req, err := http.NewRequest("POST", "/questions", bytes.NewBuffer(jsonText))
	require.NoError(t, err, "Failed to create request")

	rr := httptest.NewRecorder()
	handler.CreateQuestion(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code, "wrong status")
}
