package port

import (
	"api/internal/entity/domain"
	"api/internal/entity/model"
)

type AnswerService interface {
	SaveAnswer(answer *[]domain.AnswerChoice) error
	GetAnswerListbyQuestionID(id int) ([]model.Answer, error)
	GetAnswerCorrectionbyID(id int) (bool, error)
}

type AnswerRepository interface {
	Save(answer *[]domain.AnswerChoice) error
	GetAnswerListbyQuestionID(id int) ([]domain.AnswerChoice, error)
	GetAnswerCorrectionbyID(id int) (bool, error)
}
