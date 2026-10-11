package handler

import (
	"api/internal/entity/model"
	"api/internal/port"

	"github.com/gin-gonic/gin"
)

type QuestionHandler struct {
	QuestionService port.QuestionService
}

func NewQuestionHandler(QuestionService port.QuestionService) *QuestionHandler {
	return &QuestionHandler{
		QuestionService: QuestionService,
	}
}

func (h *QuestionHandler) SaveQuestion(c *gin.Context) {
	var questions []model.Question
	if err := c.ShouldBindJSON(&questions); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.QuestionService.SaveQuestion(questions)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Questions saved successfully"})
}

func (h *QuestionHandler) GetAllQuestions(c *gin.Context) {
	questions, err := h.QuestionService.GetAllQuestions()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, questions)
}

func (h *QuestionHandler) GetQuestionsbyIDList(c *gin.Context) {
	var idList model.QuestionIDList
	if err := c.ShouldBindJSON(&idList); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	questions, err := h.QuestionService.GetQuestionsbyIDList(idList)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, questions)
}
