package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AdminAuthentication(c *gin.Context) {
	tokenstring := c.Request.Header.Get("autharization")
	if len(tokenstring) == 0 {
		err := errors.New("autharization header not provided")
		c.AbortWithStatusJSON(http.StatusUnauthorized,err.Error())
		return
	}
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SUPER_SECRET_KEY")), nil
	})

	if err != nil {
		c.JSON(500, gin.H{
			"error": "error occurse while token generation",
		})
		c.Abort()
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	} else {
		c.AbortWithStatus(401)
	}
}