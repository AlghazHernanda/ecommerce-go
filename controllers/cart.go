package controllers

import (
	"errors"
	"log"
	"net/http"
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
	}
}

func (app *Application) RemoveItem() gin.HandlerFunc {
}

func GetItemFromCart() gin.HandlerFunc {

}

func (app *Application) BuyFromCart() gin.HandlerFunc {
}

func (app *Application) InstantBuy() gin.HandlerFunc {
}
