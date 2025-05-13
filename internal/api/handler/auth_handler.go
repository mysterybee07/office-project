package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mysterybee07/office-project-setup/internal/api/request"
	"github.com/mysterybee07/office-project-setup/internal/api/response"
	"github.com/mysterybee07/office-project-setup/internal/auth"
	"github.com/mysterybee07/office-project-setup/internal/model"
	"gorm.io/gorm"
)

// @Summary Login user
// @Description Login with credentials
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body request.LoginRequest true "User credentials"
// @Success 200 {object} response.LoginResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/login [post]
func Login(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var credentials request.LoginRequest

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

		auth.SetToken(ctx, accessToken, refreshToken)

		// Save refresh token in database
		tokenRecord := model.Auth{
			UserId: user.ID,
			Token:  refreshToken,
		}
		if err := db.Create(&tokenRecord).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store refresh token"})
			return
		}

		ctx.JSON(http.StatusOK, response.LoginResponse{
			Email:        user.Email,
			Role:         user.Role,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			Message:      "Login successful",
		})
	}
}
