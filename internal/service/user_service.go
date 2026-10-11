package service

import (
	"api/internal/entity/domain"
	"api/internal/entity/model"
	"api/internal/port"

	"github.com/sirupsen/logrus"
)

type UserService struct {
	UserRepo      port.UserRepository
	AnswerService port.AnswerService
}

func NewUserService(userRepo port.UserRepository,
	answerService port.AnswerService) *UserService {
	return &UserService{
		UserRepo:      userRepo,
		AnswerService: answerService,
	}
}

func (s *UserService) GetUserInfo(userID string) (string, error) {
	return s.UserRepo.GetUserInfo(userID)
}

func (s *UserService) SaveUserInfo(user model.User) (*domain.User, error) {
	score, fullScore := s.CalculateScore(user.AnswerList)
	userInfo := domain.User{
		FullName:  user.FullName,
		Score:     score,
		FullScore: fullScore,
	}

	_, err := s.UserRepo.Save(&userInfo)
	if err != nil {
		logrus.Errorf("error happened while saving user info: %v", err)
		return nil, err
	}

	return &userInfo, nil
}

func (s *UserService) CalculateScore(answers []model.Answer) (int, int) {
	// TODO: validate answer id must not duplicate
	score := 0
	fullScore := 0
	for _, answer := range answers {
		isCorrect, err := s.AnswerService.GetAnswerCorrectionbyID(answer.ID)
		if err != nil {
			logrus.Errorf("error happened while calculating score: %v", err)
			//return score, err

			// Skip the answer if there's an error for now
			// continue calculating the score for other answers
			// TODO: consider returning an error or logging it for further investigation
			continue
		}

		if isCorrect {
			score++
		}
		fullScore++
	}

	return score, fullScore
}
