package controllers

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
}

func (app *Application) AddToCart() gin.Handler {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Println("product id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return

		}

		userQueryID := c.Query("user_id")
		if userQueryID == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeOut(context.Background(), 5*time.Second)

		defer cancel()

		//acces databse folder cart.go file
		err = databse.AddProductToCart(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServer, err)
		}
		c.IndentedJSON(200, "succcesfully added to cart")
	}
}

func (app *Application) RemoveItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Println("product id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return

		}

		userQueryID := c.Query("user_id")
		if userQueryID == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeOut(context.Background(), 5*time.Second)

		defer cancel()

		err = database.RemoveCartItem(ctx, app.prodCollection, app.UserCollection, ProductID, userQueryID)

		if err != nil {
			c.IndentedJSON(htttp.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(200, "successfully removes item from cart")
	}
}

func GetItemFromCart() gin.HandlerFunc {
	return func(c *gin.Context){
		user_id := c.Query("id")

		if user_id == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error":"invalid id"})
			c.Abort()
			return
		
		}	

		usert_id, _:= primitive.ObjectIDFormatHex(user_id)

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var filledcart models.user
		err := UserCollection.FindOne(ctx, bson.D{primitive.E{Key:"_id", Value:usert_id}}).Decode(&filledcart)
	
		if err != nil {
			log.Println(err)
			c.IndentedJSON(500, "not found")
			return
		}
		
		filter_match := bson.D{{Key:$"match", Value:bson.D{primitive.E{Key:"_id",Value:usert_id}}}}
		unwind := bson.D{{Key:"$unwind", Value: bson.D{primitive.E{Key:"path", Value:"$usercart"}}}}
		grouping := bson.D{{Key:"$group", Value:bson.D{primitive.E{Key:"_id", Value:"$_id"}, {Key: "total",Value:bson.D{primitive.E{Key:"$sum", Value:"$usercart.price"}}}}}}
		pointcursor, err := UserCollection.aggregate(ctx, mongo.Pipeline{filter_match,unwind,grouping})
		
		usert_id, _:= primitive.ObjectIDFormatHex(user_id)

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var filledcart models.user
		err := UserCollection.FindOne(ctx, bson.D{primitive.E{Key:"_id", Value:usert_id}}).Decode(&filledcart)
	
		if err != nil {
			log.Println(err)
			c.IndentedJSON(500, "not found")
			return
		}
		
		filter_match := bson.D{{Key:$"match", Value:bson.D{primitive.E{Key:"_id",Value:usert_id}}}}
		unwind := bson.D{{Key:"$unwind", Value: bson.D{primitive.E{Key:"path", Value:"$usercart"}}}}
		grouping := bson.D{{Key:"$group", Value:bson.D{primitive.E{Key:"_id", Value:"$_id"}, {Key: "total",Value:bson.D{primitive.E{Key:"$sum", Value:"$usercart.price"}}}}}}
		pointcursor, err := UserCollection.aggregate(ctx, mongo.Pipeline{filter_match,unwind,grouping})
		
		if err != nil{
			log.Println(err)
		}	
		
		var listing []bson.M
		if err = pointcursor.All(ctx, &listing); err!=nil{
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		for _, json := range listing {
			c.IndentedJSON(200, json["total"])
			c.IndentedJSON(200, filledcart.UserCart)

		}
		ctx.Done()
		}	
		}	
}

func (app *Application) BuyFromCart() gin.HandlerFunc {
	return func(c *Application){
		userQueryID := c.Query("id")
		if userQueryID == ""{
			log.Panicln("user id is empty")
			c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		defer cancel()

		err := database.BuyItemFromCart(ctx, app.Collection, userQueryID)
		if err! = nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			c.IndentedJSON("successfully added item to cart")
		}
	}
}

func (app *Application) InstantBuy() gin.HandlerFunc {

	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		if productQueryID == "" {
			log.Println("product id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return

		}

		userQueryID := c.Query("user_id")
		if userQueryID == "" {
			log.Println("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeOut(context.Background(), 5*time.Second)

		defer cancel()

		err = database.InstantBuyer(ctx, app.prodCollection, app.userCollection, productID, userQueryID)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "successfully ORDER placed")
	}

}
