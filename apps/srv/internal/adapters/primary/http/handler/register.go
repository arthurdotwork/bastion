package handler

import (
	"log/slog"
	"net/http"

	"github.com/arthurdotwork/bastion/internal/domain/membership"
	"github.com/gin-gonic/gin"
)

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(rs *membership.RegisterService) gin.HandlerFunc {
	return func(context *gin.Context) {
		ctx := context.Request.Context()

		var req registerRequest
		if err := context.ShouldBindJSON(&req); err != nil {
			slog.ErrorContext(ctx, "failed to validate request", slog.Any("error", err))
			context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		if _, err := rs.Register(ctx, membership.User{Email: req.Email, Password: req.Password}); err != nil {
			slog.ErrorContext(ctx, "failed to register user", slog.Any("error", err))
			context.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		context.JSON(http.StatusCreated, gin.H{})
	}
}
