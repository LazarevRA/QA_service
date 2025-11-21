package models

import "time"

//Модель ответа
type Answer struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	QuestionID int       `gorm:"not null" json:"question_id"`
	UserID     string    `gorm:"type:uuid;not null" json:"user_id"`
	Text       string    `gorm:"type:text;not null" json:"text"`
	CreatedAt  time.Time `json:"created_at"`
}
