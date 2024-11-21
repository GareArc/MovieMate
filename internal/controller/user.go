package controller

import (
	"net/http"

	"github.com/GareArc/MovieMate/internal/service"
	"github.com/GareArc/MovieMate/internal/types/model"
	"github.com/gin-gonic/gin"
)

type UserAuthBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname"`
}

type UserReturnBody struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type UserController struct{}

func (uc *UserController) CheckActiveUser(c *gin.Context) {
	// assume user is in context
	user := c.MustGet("current_user").(model.User)
	c.JSON(200, gin.H{
		"user": &UserReturnBody{
			ID:       user.ID,
			Email:    user.Email,
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
		},
	})
}

func (uc *UserController) LoginUser(c *gin.Context) {
	var req UserAuthBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user_service := service.AuthService{}
	jwt_token, user, err := user_service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"jwt_token": jwt_token,
		"user": &UserReturnBody{
			ID:       user.ID,
			Email:    user.Email,
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
		},
	})

}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var req UserAuthBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Nickname == "" {
		req.Nickname = "TBD"
	}

	user_service := service.AuthService{}
	jwt_token, user, err := user_service.Register(req.Email, req.Password, req.Nickname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"jwt_token": jwt_token,
		"user": &UserReturnBody{
			ID:       user.ID,
			Email:    user.Email,
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
		},
	})

}
