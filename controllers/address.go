package controllers

import (
	"net/http"

	"github.com/AcevedoEsteban/goEcommerce-yt/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddAddresss() gin.HandlerFunc {

}

func EditHomeAddress() gin.HandlerFunc {

}

func EditWorkAddress() gin.HandlerFunc {

}

func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id := c.Query("id")

		if user_id == "" {
			c.Header("Context-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid Search Index"})
			c.Abort()
			return
		}

		addresses := make([]models.Address, 0)

		usert_id, err := primitive.ObjectIDFromHex(user_id)

		if err != nil {
			c.IndentedJSON(500, "InternaL Server Error")
		}

	}

}
