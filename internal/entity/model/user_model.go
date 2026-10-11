package model

type User struct {
	FullName   string   `json:"name"`
	AnswerList []Answer `json:"answer_list"`
}
