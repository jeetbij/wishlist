package users

import (
	"example/bucket/app/helpers"
	"example/bucket/app/models/user"

	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(ctx *gin.Context) {
	body := helpers.SignUpRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	usr := user.User{UserName: body.UserName, Email: body.Email, Password: string(hash)}
	result := user.DB().Create(&usr)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create the user",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": "User created successfully",
	})
}

func LogIn(ctx *gin.Context) {
	body := helpers.LogInRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	var usr user.User
	user.DB().First(&usr, "email = ?", body.Email)
	if usr.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(body.Password))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": usr.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate a token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func Validate(ctx *gin.Context) {
	_, notGuest := ctx.Get("user")

	if notGuest {
		ctx.JSON(http.StatusOK, gin.H{
			"success": "I am a logged in user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": "I am a guest user",
	})
}
