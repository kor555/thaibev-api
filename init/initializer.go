package initializer

import (
	"api/config"
	"api/internal/handler"
	"api/internal/repo"
	"api/internal/service"
	"fmt"
)

type Initializer struct {
	HealthHandler *handler.HttpHealthHandler

	QuestionHandler *handler.QuestionHandler
	UserInfoHandler *handler.UserInfoHandler
}

func Init(cfg config.Config) *Initializer {

	database, err := connectToDatabase(cfg)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// Health Check
	HealthHandler := handler.NewHealthHandler()

	//repo
	answerRepo := repo.NewAnswerRepository(database)
	questionRepo := repo.NewQuestionRepository(database)
	userRepo := repo.NewUserRepository(database)

	//service
	answerService := service.NewAnswerService(answerRepo)
	questionService := service.NewQuestionService(questionRepo, answerService)
	userService := service.NewUserService(userRepo, answerService)

	//handler
	questionHandler := handler.NewQuestionHandler(questionService)
	userInfoHandler := handler.NewUserInfoHandler(userService)

	return &Initializer{
		HealthHandler: HealthHandler,

		QuestionHandler: questionHandler,
		UserInfoHandler: userInfoHandler,
	}
}
