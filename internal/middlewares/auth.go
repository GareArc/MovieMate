package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/GareArc/MovieMate/internal/db"
	"github.com/GareArc/MovieMate/internal/types/model"
	"github.com/GareArc/MovieMate/internal/utils"
)

func RequireLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth_string := c.GetHeader("Authorization")
		if auth_string == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		auth_token := strings.Split(auth_string, " ")
		if len(auth_token) != 2 || auth_token[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "Authorization header must be in the format `Bearer <token>`"})
			c.Abort()
			return
		}

		claims, ok := utils.VerifyJWTToken(auth_token[1])
		if ok != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		user_id := claims["user_id"].(string)
		// get user from db by id
		var user model.User
		db.MainDB.Model(&model.User{}).Where("id = ?", user_id).First(&user)

		// if user not found
		if user.ID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		c.Set("current_user", user)
		c.Next()
	}
}
