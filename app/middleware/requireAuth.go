package middleware

import (
	"example/bucket/app/models/user"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(ctx *gin.Context) {
	token := ctx.Request.Header.Get("X-Token")

	if token != "" {
		//validate token if present
		token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			var usr user.User
			user.DB().First(&usr, "email = ?", claims["email"])
			if usr.ID == 0 {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			ctx.Set("user", usr)
		} else {
			fmt.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

	}

	ctx.Next()
}
