package port

import (
	"api/internal/entity/domain"
	"api/internal/entity/model"
)

type QuestionService interface {
	SaveQuestion(questions []model.Question) error
	GetQuestionsbyIDList(idList model.QuestionIDList) (*[]model.Question, error)
	GetAllQuestions() (*[]model.Question, error)
}

type QuestionRepository interface {
	Save(question *domain.Question) (*domain.Question, error)
	GetQuestionbyID(id int) (*domain.Question, error)
	GetAllQuestionID() (*[]domain.Question, error)
}
