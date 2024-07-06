package userService

import (
	"context"
	"net/http"
	"takedi/xApi/infra/database"
	userSchema "takedi/xApi/infra/schemas"
	"takedi/xApi/utils/hashes"
	responseJson "takedi/xApi/utils/responses"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()

var user userSchema.User
var userCollection *mongo.Collection = database.GetCollection(
	database.Client,
	"users",
)

func InsertOne(c *gin.Context, ctx context.Context) {
	// Validate the request body
	err := c.BindJSON(&user)

	if err != nil {
		responseJson.Error(
			c,
			http.StatusBadRequest,
			"Error",
			err,
		)

		return
	}

	// Use the validator library to validate required fields
	validationErr := validate.Struct(&user)

	if validationErr != nil {
		responseJson.Error(
			c,
			http.StatusBadRequest,
			"Error",
			validationErr,
		)

		return
	}

	hashedPassword, err := hashes.HashPassword(user.Password)

	if err != nil {
		responseJson.Error(
			c,
			http.StatusBadRequest,
			"Error",
			err,
		)

		return
	}

	newUserData := userSchema.User{
		Name:       user.Name,
		Phone:      user.Phone,
		Birthday:   user.Birthday,
		Password:   string(hashedPassword),
		Created_At: time.Now(),
		Updated_At: time.Now(),
		Is_Deleted: false,
	}

	_, err = userCollection.InsertOne(ctx, newUserData)

	if err != nil {
		responseJson.Error(
			c,
			http.StatusInternalServerError,
			"Error",
			err,
		)

		return
	}

	responseJson.Text(
		c,
		http.StatusCreated,
		"User created successfully!",
	)
}

func FindOneById(c *gin.Context, ctx context.Context) {
	id := c.Param("userId")

	objectId, _ := primitive.ObjectIDFromHex(id)
	err := userCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&user)

	if err != nil {
		responseJson.Error(
			c,
			http.StatusInternalServerError,
			"Error",
			err,
		)

		return
	}

	responseJson.Data(
		c,
		http.StatusOK,
		"User searched successfully!",
		map[string]interface{}{
			"name":      user.Name,
			"birthday":  user.Birthday,
			"createdAt": user.Created_At,
		},
	)
}

func UpdateOne(c *gin.Context, ctx context.Context) {
	id := c.Param("userId")

	// Validate the request body
	bindErr := c.BindJSON(&user)
	if bindErr != nil {
		responseJson.Error(
			c,
			http.StatusBadRequest,
			"Error",
			bindErr,
		)

		return
	}

	// Use the validator library to validate required fields
	validationErr := validate.Struct(&user)
	if validationErr != nil {
		responseJson.Error(
			c,
			http.StatusBadRequest,
			"Error",
			validationErr,
		)

		return
	}

	dataToUpdate := bson.M{}
	dataToUpdate["updatedAt"] = time.Now()

	if user.Name != "" {
		dataToUpdate["name"] = user.Name
	}

	if user.Phone != "" {
		dataToUpdate["phone"] = user.Phone
	}

	if user.Birthday != "" {
		dataToUpdate["birthday"] = user.Birthday
	}

	objectId, _ := primitive.ObjectIDFromHex(id)
	_, err := userCollection.UpdateOne(
		ctx,
		bson.M{"_id": objectId},
		bson.M{"$set": dataToUpdate},
	)

	if err != nil {
		responseJson.Error(
			c,
			http.StatusInternalServerError,
			"Error",
			err,
		)

		return
	}

	responseJson.Text(
		c,
		http.StatusOK,
		"User updated successfully!",
	)
}

func DeleteOne(c *gin.Context, ctx context.Context) {
	id := c.Param("userId")

	objectId, _ := primitive.ObjectIDFromHex(id)
	_, err := userCollection.DeleteOne(ctx, bson.M{"_id": objectId})

	if err != nil {
		responseJson.Error(
			c,
			http.StatusInternalServerError,
			"Error to delete user.",
			err,
		)

		return
	}

	responseJson.Text(
		c,
		http.StatusOK,
		"User deleted successfully!",
	)
}
