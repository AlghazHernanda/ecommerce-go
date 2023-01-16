package token

import (
	"log"
	"os"
	"time"

	"github.com/AlghazHernanda/ecommerce-go/database"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type signedDetails struct {
	Email      string `json:"email"`
	First_Name string
	Last_Name  string
	Uid        string
	jwt.StandardClaims
}

var UserData *mongo.Collection = database.UserData(database.Client, "Users")

var SECRETE_KEY = os.Getenv("SECRETE_KEY")

func TokenGenerator(email string, firstname string, lastname string, uid string) (signedtoken string, signedrefreshtoken string, err error) {
	claims := &signedDetails{
		Email:      email,
		First_Name: firstname,
		Last_Name:  lastname,
		Uid:        uid,
		StandardClaims: jwt.StandardClaims{
			ExpireAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshclaims := &signedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpireAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRETE_KEY))

	if err != nil {
		return "", "", err

	}

	refreshtoken, err := jwt.NewWithClaims(jwt.SigningMethodHS384, refreshclaims).SignedString([]byte(SECRETE_KEY))
	if err != nil {
		log.Panic(err)
		return
	}
	return token, refreshtoken, err
}
