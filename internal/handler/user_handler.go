package handler

import (
	"api/internal/entity/model"
	"api/internal/port"

	"github.com/gin-gonic/gin"
)

type UserInfoHandler struct {
	UserService port.UserService
}

func NewUserInfoHandler(userService port.UserService) *UserInfoHandler {
	return &UserInfoHandler{
		UserService: userService,
	}
}

func (h *UserInfoHandler) GetUserInfo(c *gin.Context) {
	userID := c.Param("id")
	userInfo, err := h.UserService.GetUserInfo(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, userInfo)
}

func (h *UserInfoHandler) SaveUserInfo(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userInfo, err := h.UserService.SaveUserInfo(user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User info saved successfully", "user": userInfo})
}
