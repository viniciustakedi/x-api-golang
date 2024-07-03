package controllers

import (
	"context"
	"net/http"
	"takedi/xApi/configs"
	"takedi/xApi/models"
	"takedi/xApi/responses"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		var user models.User
		defer cancel()

		// Validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(
				http.StatusBadRequest,
				responses.UserResponse{
					Status:  http.StatusBadRequest,
					Message: "Error",
					Data: map[string]interface{}{
						"data": err.Error(),
					},
				},
			)

			return
		}

		// Use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				responses.UserResponse{
					Status:  http.StatusBadRequest,
					Message: "error",
					Data: map[string]interface{}{
						"data": validationErr.Error(),
					},
				},
			)

			return
		}

		newUser := models.User{
			Name:     user.Name,
			Location: user.Location,
			Title:    user.Title,
		}

		result, err := userCollection.InsertOne(ctx, newUser)

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				responses.UserResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data: map[string]interface{}{
						"data": err.Error(),
					},
				},
			)

			return
		}

		c.JSON(
			http.StatusCreated,
			responses.UserResponse{
				Status:  http.StatusCreated,
				Message: "Success",
				Data: map[string]interface{}{
					"InsertedID": result.InsertedID,
				},
			},
		)
	}
}

func GetUserById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		var user models.User
		defer cancel()

		userId := c.Param("userId")
		objId, _ := primitive.ObjectIDFromHex(userId)

		err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				responses.UserResponse{
					Status:  http.StatusInternalServerError,
					Message: "Error",
					Data: map[string]interface{}{
						"data": err.Error(),
					},
				},
			)

			return
		}

		c.JSON(
			http.StatusOK,
			responses.UserResponse{
				Status:  http.StatusOK,
				Message: "Success",
				Data: map[string]interface{}{
					"name":     user.Name,
					"location": user.Location,
					"title":    user.Title,
				},
			},
		)
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		var user models.User
		defer cancel()

		userId := c.Param("userId")
		objId, _ := primitive.ObjectIDFromHex(userId)

		// Validate the request body
		bindErr := c.BindJSON(&user)

		if bindErr != nil {
			c.JSON(
				http.StatusBadRequest,
				responses.UserResponse{
					Status:  http.StatusBadRequest,
					Message: "error",
					Data: map[string]interface{}{
						"data": bindErr.Error(),
					},
				},
			)

			return
		}

		// Use the validator library to validate required fields
		validationErr := validate.Struct(&user)

		if validationErr != nil {
			c.JSON(
				http.StatusBadRequest,
				responses.UserResponse{
					Status:  http.StatusBadRequest,
					Message: "error",
					Data: map[string]interface{}{
						"data": validationErr.Error(),
					},
				},
			)

			return
		}

		dataToUpdate := bson.M{
			"name":     user.Name,
			"location": user.Location,
			"title":    user.Title,
		}

		result, err := userCollection.UpdateOne(
			ctx,
			bson.M{"_id": objId},
			bson.M{"$set": dataToUpdate},
		)

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				responses.UserResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data: map[string]interface{}{
						"data": err.Error(),
					},
				},
			)

			return
		}

		c.JSON(
			http.StatusOK,
			responses.UserResponse{
				Status:  http.StatusOK,
				Message: "success",
				Data: map[string]interface{}{
					"modifiedCount": result.ModifiedCount,
				},
			},
		)
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		userId := c.Param("userId")
		objId, _ := primitive.ObjectIDFromHex(userId)

		result, err := userCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				responses.UserResponse{
					Status:  http.StatusInternalServerError,
					Message: "Error",
					Data: map[string]interface{}{
						"data": err.Error(),
					},
				},
			)

			return
		}

		c.JSON(
			http.StatusOK,
			responses.UserResponse{
				Status:  http.StatusOK,
				Message: "Success",
				Data: map[string]interface{}{
					"deletedCount": result.DeletedCount,
				},
			},
		)
	}
}
