package middleware

import (
	"MyGram/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationMiddleware interface {
	Authenticate() gin.HandlerFunc
}

type Middleware struct {
	tokenRepository repositories.TokenRepository
}

func NewMiddleware(tokenRepository repositories.TokenRepository) *Middleware {
	return &Middleware{tokenRepository: tokenRepository}
}

func (m *Middleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		user, err := m.tokenRepository.FindUserByToken(authHeader)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
