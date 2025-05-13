package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mysterybee07/office-project-setup/internal/auth"
	"github.com/mysterybee07/office-project-setup/internal/model"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var credentials struct {
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required"`
		}

		if err := ctx.ShouldBindJSON(&credentials); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request format: " + err.Error(),
			})
			return
		}
		fmt.Println(credentials)

		var user model.User

		if err := db.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid email or password",
			})
			return
		}

		// Compare hashed password
		// if !auth.CheckPasswordHash(user.Password, credentials.Password) {
		// 	ctx.JSON(http.StatusUnauthorized, gin.H{
		// 		"error": "Invalid email or password",
		// 	})
		// 	return
		// }

		accessToken, refreshToken, err := auth.GenerateJWTToken(user.Email, user.Role)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		// Save refresh token in database
		tokenRecord := model.Auth{
			UserId: user.ID,
			Token:  refreshToken,
		}
		if err := db.Create(&tokenRecord).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store refresh token"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"email":         user.Email,
			"role":          user.Role,
			"access_token":  accessToken,
			"refresh_token": refreshToken,
			"message":       "Login successful",
		})
	}
}
