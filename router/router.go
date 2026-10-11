package router

import (
	initializer "api/init"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Router struct {
	Initializer *initializer.Initializer
}

func NewRouter(Initializer *initializer.Initializer) *Router {
	return &Router{
		Initializer: Initializer,
	}
}

func (r *Router) CreateRouter() *gin.Engine {
	router := gin.New()
	router.Use(LoggingMiddleware())
	router.Use(gin.Recovery())

	api := router.Group("/api")

	r.setupHealthCheck(api)

	r.setupQuestion(api)
	r.setupUserInfo(api)

	return router
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		logrus.WithFields(logrus.Fields{
			"method":   c.Request.Method,
			"path":     c.Request.URL.Path,
			"status":   c.Writer.Status(),
			"duration": time.Since(start).String(),
		}).Info("handled request")
	}
}

func (r *Router) setupHealthCheck(api *gin.RouterGroup) {
	health := api.Group("/health")

	health.GET("", r.Initializer.HealthHandler.Healthcheck)
	health.GET("/version", r.Initializer.HealthHandler.Version)
}

func (r *Router) setupQuestion(api *gin.RouterGroup) {
	question := api.Group("/question")
	question.POST("/save", r.Initializer.QuestionHandler.SaveQuestion)
	question.GET("", r.Initializer.QuestionHandler.GetAllQuestions)
	question.POST("/getQuestionsbyIDList", r.Initializer.QuestionHandler.GetQuestionsbyIDList)
}

func (r *Router) setupUserInfo(api *gin.RouterGroup) {
	user := api.Group("/user")
	user.GET("/:id", r.Initializer.UserInfoHandler.GetUserInfo)
	user.POST("/save", r.Initializer.UserInfoHandler.SaveUserInfo)
}
