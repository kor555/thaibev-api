package service

import (
	"api/internal/entity/domain"
	"api/internal/entity/model"
	"api/internal/port"
)

type QuestionService struct {
	QuestionRepo  port.QuestionRepository
	AnswerService port.AnswerService
}

func NewQuestionService(QuestionRepo port.QuestionRepository,
	AnswerService port.AnswerService) *QuestionService {
	return &QuestionService{
		QuestionRepo:  QuestionRepo,
		AnswerService: AnswerService,
	}
}

func (s *QuestionService) SaveQuestion(questions []model.Question) error {
	for _, question := range questions {
		// TODO: validate answer must have at only one correct answer
		var q domain.Question
		q.ID = uint(question.ID)
		q.QuestionText = question.QuestionText

		_, err := s.QuestionRepo.Save(&q)
		if err != nil {
			return err
		}

		answers := prepareAnswersList(question.Answers, q.ID)
		err = s.AnswerService.SaveAnswer(&answers)
		if err != nil {
			return err
		}
	}

	return nil
}

func prepareAnswersList(answers []model.Answer, questionID uint) []domain.AnswerChoice {
	answerList := []domain.AnswerChoice{}
	for _, answer := range answers {
		answerList = append(answerList, domain.AnswerChoice{
			ID:         uint(answer.ID),
			ChoiceText: answer.ChoiceText,
			QuestionID: questionID,
			IsCorrect:  answer.IsCorrect,
		})
	}
	return answerList
}

func (s *QuestionService) GetAllQuestions() (*[]model.Question, error) {
	idList, err := s.QuestionRepo.GetAllQuestionID()
	if err != nil {
		return nil, err
	}

	questionList, err := s.GetQuestionsbyIDList(prepareIDList(idList))
	if err != nil {
		return nil, err
	}

	return questionList, nil
}

func prepareIDList(idList *[]domain.Question) model.QuestionIDList {
	var list model.QuestionIDList
	for _, i := range *idList {
		list.ID = append(list.ID, int(i.ID))
	}

	return list
}

func (s *QuestionService) GetQuestionsbyIDList(idList model.QuestionIDList) (*[]model.Question, error) {
	questionList := []model.Question{}

	for _, id := range idList.ID {
		q, err := s.QuestionRepo.GetQuestionbyID(id)
		if err != nil {
			return nil, err
		}

		if q.ID == 0 {
			continue // Skip if question not found
		}

		answers, err := s.AnswerService.GetAnswerListbyQuestionID(id)
		if err != nil {
			return nil, err
		}

		question := model.Question{
			ID:           int(q.ID),
			QuestionText: q.QuestionText,
			Answers:      answers,
		}
		questionList = append(questionList, question)
	}

	return &questionList, nil
}
