package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {

}

func SignUp() gin.HandlerFunc {

}

func Login() gin.HandlerFunc {

}

func ProductViewerAdmin() gin.HandlerFunc {

}
