package service

import (
	"api/internal/entity/domain"
	"api/internal/entity/model"
	"api/internal/port"
)

type AnswerService struct {
	AnswerRepo port.AnswerRepository
}

func NewAnswerService(AnswerRepo port.AnswerRepository) *AnswerService {
	return &AnswerService{
		AnswerRepo: AnswerRepo,
	}
}

func (s *AnswerService) SaveAnswer(answer *[]domain.AnswerChoice) error {
	return s.AnswerRepo.Save(answer)

}

func (s *AnswerService) GetAnswerListbyQuestionID(id int) ([]model.Answer, error) {
	list, err := s.AnswerRepo.GetAnswerListbyQuestionID(id)
	if err != nil {
		return nil, err
	}

	answers := []model.Answer{}
	// loop to flag out the correct answer
	for _, l := range list {
		if l.IsCorrect {
			l.IsCorrect = false
		}

		answers = append(answers, model.Answer{
			ID:         int(l.ID),
			ChoiceText: l.ChoiceText,
			IsCorrect:  l.IsCorrect,
		})
	}

	return answers, nil
}

func (s *AnswerService) GetAnswerCorrectionbyID(id int) (bool, error) {
	isCorrect, err := s.AnswerRepo.GetAnswerCorrectionbyID(id)
	if err != nil {
		return false, err
	}

	return isCorrect, nil
}
