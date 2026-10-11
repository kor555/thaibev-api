package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName  string `gorm:"column:name" json:"name"`
	Score     int    `gorm:"column:score" json:"score"`
	FullScore int    `gorm:"column:full_score" json:"full_score"`
}
