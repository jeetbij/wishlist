package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"example/bucket/app/helpers"
	"log"

	"github.com/gin-gonic/gin"
)

func generateToken() string {
	length := 32
	b := make([]byte, length)
	_, err := rand.Read(b)
	for err != nil {
		log.Println(err)
		return generateToken()
	}
	return hex.EncodeToString(b)
}

func SetGuestToken(ctx *gin.Context) {
	token, err := ctx.Cookie("token")
	if err != nil {
		log.Println(err)
		token = generateToken()
		helpers.SetCookie(ctx, "token", token, (3600 * 24 * 30 * 12))
	}
	ctx.Set("token", token)
	ctx.Next()
}
