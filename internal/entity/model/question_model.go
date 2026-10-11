package model

type Question struct {
	ID           int      `json:"id"`
	QuestionText string   `json:"question"`
	Answers      []Answer `json:"choices"`
}

type QuestionIDList struct {
	ID []int `json:"id"`
}
