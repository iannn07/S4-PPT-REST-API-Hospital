package Middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"HospitalFinpro/Helper"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isLogout := c.GetBool("isLogout")
		if isLogout {
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := map[string]string{"message": "Unauthorized"}
			Helper.ResponseJSON(c.Writer, http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			response := map[string]string{"message": "Unauthorized"}
			Helper.ResponseJSON(c.Writer, http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		tokenString := authHeaderParts[1]
		// Mengecek apakah token ada dalam daftar token yang valid
		if _, ok := ValidTokens[tokenString]; !ok {
			response := map[string]string{"message": "Unauthorized"}
			Helper.ResponseJSON(c.Writer, http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("asdffdasasdffasd"), nil // Ganti "your-secret-key" dengan secret key yang sebenarnya
		})

		if err != nil || !token.Valid {
			response := map[string]string{"message": "Unauthorized"}
			Helper.ResponseJSON(c.Writer, http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)
		c.Set("userID", claims["userID"])
		c.Next()
	}
}