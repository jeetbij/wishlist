package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"

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
		// Need to change domain name later
		ctx.SetSameSite(http.SameSiteLaxMode)
		ctx.SetCookie("token", token, (3600 * 24 * 30 * 12), "", "", false, true)
	}
	ctx.Set("token", token)
	ctx.Next()
}
