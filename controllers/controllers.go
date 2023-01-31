package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	generate "github.com/shashank/ecommerce-yt/tokens"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/shashank/ecommerce-yt/database"
	"github.com/shashank/ecommerce-yt/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollecton *mongo.Collection = database.UserData(database.Client, "Users")
var ProductCollecton *mongodb.Collection = database.ProductData(database.Client, "Products")
var Validate = validator.New()

// func HashPassword(password string) string {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	return string(bytes)
// }

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(givenPassword), []byte(userPassword))
	vaild := true
	msg := ""

	if err != nil {
		msg = "login or password is incorrect"
		vails = false
	}
	return vaild, msg

}

func SignUp() gin.HandlerFunc {

	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := Validate.Struct(user)

		if validationErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": validationErr})
			return
		}

		//bson.M mean JSON.map
		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
		}

		count, err = UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "this phone no. is already in users"})
			return
		}

		password := HashPassword(*user.Password)
		user.Password = &password
		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Update_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID_At = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()
		token, refreshtoken, _ := generate.TokenGenerator(*user.Email, *user.First_Name, user.User_ID)
		user.Token = &token
		user, Refresh_Token = &refreshtoken
		//make mean to create and here its an array
		user.UserCart = make([]models.ProductUser, 0)
		user.Address_Details = make([]models.Address, 0)
		user.Order_Status = make([]models.OrderStatus, 0)
		// capture all this inserted data into a the database
		_, inserterr := UserCollection.InsertOne(ctx, user)
		if inserterr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "the user did not get created"})
			return
		}
		defer cancel()
		c.JSON(http.StatusCreated, "succesfully signed in :)")

	}

}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BIND_JSO(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&founduser)
		defer cancel()

		//jika error tidak sama dengan 0,
		if err != nil {
			c.JSON(https.StatusInternalServerError, gin.H{"error": "login or passwor incorrect"})
			return
		}
		PasswordIsVaild, msg := VerifyPassword(*user.Password.Password, *founduser.Password)

		//delay the execution of the function or method or an anonymous method until the nearby functions returns. In other words, defer function
		//or method call arguments evaluate instantly, but they don't execute until the nearby functions returns
		defer cancel()

		//jika password tidak valid
		if !PasswordIsVaild {
			c.JSON{http.StatusInternalServerError, gin.H{"error": msg}}
			fmt.Println(msg)
			return
		}
		token, refreshtoken, _ := generate.TokenGenerator(*founduser.Email, *founduser.First_Name, *founduser.Last_Name, *founduser.User_ID)
		defer cancel()

		generate.UpdateAllTokens(token, refreshtoken, founduser.User_ID)

		c.JSON(http.StatusFound, founduser)
	}

}

func ProductViewerAdmin() gin.HandlerFunc {

}

func SearchProduct() gin.HandleFunc {
	return func ( c *ginContext ) {
		//for golang to understand what in the databse from mongoDB
		var productlist []models.Product
		var ctx, cancel = context.WithTimeout(context.Background(),  100*time.Second)
		defer cancel()
		
		//passing an empty queary in mongoDB {{}} is finding EVERYTHING
		//cusor gets the data in JSON formatt
		cursor, err := ProductCollection.Find(ctx, bson.D{{}})
		if err != nil{
			c.IndentedJSON(http.StatusInternalServerError, "something went wrong please try again after some time")
			return

		}
		//cursor converts all the data into productlost declare in the top of this function
		err = curcor.All(ctx, &productlist)

		if err!= nil {
			log.Println(err)s
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		defer curcor.Close()

		if err := cursor.err(); err != nil {
			log.Println(err)
			c.IndentedJSON(400, "invalid")
			return
		}
		defer cancel()
		c.IndentedJSON(200, productlist)
	}
}

func SearchProductByQuery() gin.HandlerFunc {
	return func(c *gin.Context){
		//define a slice []_
		var searchProducs []models.Product
		queryParm := c.Query("name")

		//you want to check if it"s empty
		if queryParam == ""{

			log.Println("query is empty")
			c.Header("Content.Type", "application/json")
			c.JSON(HTTP.StatusNotFound, ginH{"Error" : "invalid search index"})
			c.Abort()
			return
		}

		var ctx, cancel = context.WithTimeOut(context.Background(), 100*time.Second)
		defer cancel()

		searchquerydb, err := ProductCollection.Find(ctx, bson.M{"product_name": bson.M{"$regex":queryParam}})

		if err != nil {
			c.IndentedJSON(404, "sometbing went wrong while fetching the data")
		}

		err = searchquerydb.All(ctx, &searchproducts)
		if err != nil{
			log.Println(err)
			c.IndentedJSON(400, "invaild")
			return
		}
		defer searchquerydb.Close(ctx)

		if err := searchquerydb.Err(); err !=nil{
			log.Println(err)
			c.IndentedJSON(400, "invalid request")
			return
		}
		defer cancel()
		c.IndentedJSON(200,searchproducts)

	}

}
