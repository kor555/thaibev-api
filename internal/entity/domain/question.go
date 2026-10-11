package domain

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	QuestionText string `gorm:"column:question" json:"question"`
}
