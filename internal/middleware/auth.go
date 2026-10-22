package middleware

import (
	"api/internal"
	"errors"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const userIDKey = "userID"

// Auth validates the HS256 JWT in the "Authorization: Bearer <token>" header
// and stores the token's subject (the user ID) on the context.
func Auth() gin.HandlerFunc {
	secret := []byte(os.Getenv("JWT_SECRET"))
	if len(secret) == 0 {
		panic("middleware.Auth: JWT_SECRET is not set")
	}

	parser := jwt.NewParser(
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithExpirationRequired(),
	)
	keyFunc := func(*jwt.Token) (any, error) { return secret, nil }

	return func(c *gin.Context) {
		tokenString, err := bearerToken(c.GetHeader("Authorization"))
		if err != nil {
			abortUnauthorized(c, err.Error())
			return
		}

		claims := &jwt.RegisteredClaims{}
		if _, err := parser.ParseWithClaims(tokenString, claims, keyFunc); err != nil {
			abortUnauthorized(c, "invalid or expired token")
			return
		}
		if claims.Subject == "" {
			abortUnauthorized(c, "token has no subject")
			return
		}

		c.Set(userIDKey, claims.Subject)
		c.Next()
	}
}

// UserIDFromRequest returns the authenticated user's ID set by Auth.
func UserIDFromRequest(c *gin.Context) (string, bool) {
	id, ok := c.Get(userIDKey)
	if !ok {
		return "", false
	}
	s, ok := id.(string)
	return s, ok
}

func bearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("missing authorization header")
	}
	scheme, token, found := strings.Cut(header, " ")
	if !found || !strings.EqualFold(scheme, "Bearer") || strings.TrimSpace(token) == "" {
		return "", errors.New("authorization header must use the Bearer scheme")
	}
	return strings.TrimSpace(token), nil
}

func abortUnauthorized(c *gin.Context, message string) {
	apiErr := internal.UnauthorizedError(message, nil)
	c.AbortWithStatusJSON(apiErr.StatusCode, gin.H{
		"code":    apiErr.Code,
		"message": apiErr.Message,
	})
}
