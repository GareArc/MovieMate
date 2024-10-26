package user_controller

import (
	"net/http"

	"github.com/GareArc/MovieMate/internal/service"
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

func LoginUser(c *gin.Context) {
	var req UserAuthBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt_token, user, err := service.Login(req.Email, req.Password)
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

func RegisterUser(c *gin.Context) {
	var req UserAuthBody
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Nickname == "" {
		req.Nickname = "TBD"
	}

	jwt_token, user, err := service.Register(req.Email, req.Password, req.Nickname)
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
