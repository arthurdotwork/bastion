package middleware

import (
	"log/slog"
	"net/http"

	"github.com/arthurdotwork/bastion/internal/domain/authentication"
	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware(authenticationService *authentication.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		accessToken := c.GetHeader("Authorization")
		if accessToken == "" || len(accessToken) < 7 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing access token"})
			return
		}

		err := authenticationService.VerifyAccessToken(ctx, accessToken[7:])
		if err != nil {
			slog.ErrorContext(ctx, "failed to verify access token", "error", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid access token"})
			return
		}

		c.Next()
	}
}
