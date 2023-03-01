package helpers

import (
	"example/bucket/app/models/user"

	"github.com/gin-gonic/gin"
)

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
