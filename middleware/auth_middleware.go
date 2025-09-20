package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/Vis7044/GinCrud2/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, utils.ResponseHandler[string]{Status: false,Data: "Missing Authorization header"})
			c.Abort()
			return
		}
		err := godotenv.Load()
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.ResponseHandler[string]{Status: false,Data: "Missing jwtSecret header"})
			c.Abort()
			return
		}
		jwtSecret := os.Getenv("JWT_SECRET")
		tokenStr := strings.TrimPrefix(authHeader,"Bearer ")
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, utils.ResponseHandler[string]{Status: false, Data: "Invalid or expired token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			c.Set("userId", claims["userId"])
			c.Set("email", claims["email"])
		}
		c.Next()
	}
}