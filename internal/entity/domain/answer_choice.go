package domain

import "gorm.io/gorm"

type AnswerChoice struct {
	gorm.Model
	QuestionID uint   `gorm:"column:question_id" json:"question_id"`
	ChoiceText string `gorm:"column:choice_text" json:"choice_text"`
	IsCorrect  bool   `gorm:"column:is_correct" json:"is_correct"`
}
