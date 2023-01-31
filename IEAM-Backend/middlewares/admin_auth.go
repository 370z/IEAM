package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		hmacSampleSecret := []byte("IEAM_IT")
		hander := c.Request.Header.Get("Authorization")
		tokenString := strings.Replace(hander, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return hmacSampleSecret, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code": 401,
				"msg":  "authorization required",
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userId", claims["userId"])
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":  401,
				"msg":   "authorization required",
			})
			return
		}
	}
}
