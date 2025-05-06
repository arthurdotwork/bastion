package handler

import (
	"log/slog"
	"net/http"

	"github.com/arthurdotwork/bastion/internal/domain/authentication"
	"github.com/gin-gonic/gin"
)

type authenticateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Authenticate(authenticationService *authentication.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		var req authenticateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		accessToken, err := authenticationService.AuthenticateWithPassword(ctx, req.Email, req.Password)
		if err != nil {
			slog.ErrorContext(ctx, "failed to authenticate with password", "error", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"accessToken": accessToken.RawToken})
	}
}

func VerifyAuthentication(authenticationService *authentication.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		accessToken := c.GetHeader("Authorization")
		if accessToken == "" || len(accessToken) < 7 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing access token"})
			return
		}

		err := authenticationService.VerifyAccessToken(ctx, accessToken[7:])
		if err != nil {
			slog.ErrorContext(ctx, "failed to verify access token", "error", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid access token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "access token is valid"})
	}
}
