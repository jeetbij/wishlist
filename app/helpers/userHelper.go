package helpers

import (
	"example/bucket/app/models/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGuestToken(ctx *gin.Context) string {
	token, present := ctx.Get("token")
	if present {
		return token.(string)
	}
	return ""
}

func GetUser(ctx *gin.Context) (user.User, bool) {
	usr, present := ctx.Get("user")

	if present {
		return usr.(user.User), true
	}
	return user.User{}, false
}

func GetUserId(user user.User) uint {
	if user.ID == 0 {
		return 0
	}
	return user.ID
}

func SetCookie(ctx *gin.Context, name string, value string, maxAge int) {
	// Need to change domain name later
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(name, value, maxAge, "", "", false, true)
}

func DeleteCookie(ctx *gin.Context, name string) {
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(name, "", -1, "", "", false, true)
}
