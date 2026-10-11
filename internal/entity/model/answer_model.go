package model

type Answer struct {
	ID         int    `json:"id"`
	ChoiceText string `json:"choice"`
	IsCorrect  bool   `json:"is_correct"`
}
