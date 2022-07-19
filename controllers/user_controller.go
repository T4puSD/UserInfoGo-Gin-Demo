package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
	"user-info-service/db"
	"user-info-service/model"
	"user-info-service/responses"
)

var userCollection = db.GetCollection("users")
var validate = validator.New()

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user model.User
		id := c.Param("id")
		objectId, _ := primitive.ObjectIDFromHex(id)
		err := userCollection.FindOne(ctx, bson.M{"id": objectId}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}
		c.JSON(http.StatusOK, responses.UserResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": user},
		})
	}
}

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var user model.User
		// validation of request json body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}
		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newUser := model.User{
			Id:    primitive.NewObjectID(),
			Name:  user.Name,
			Email: user.Email,
		}

		result, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		c.JSON(http.StatusCreated, responses.UserResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    map[string]interface{}{"data": result},
		})
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		// validate request body
		var user model.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		// validate with the validator library
		if err := validate.Struct(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		objectId, _ := primitive.ObjectIDFromHex(id)

		update := bson.M{"name": user.Name, "email": user.Email}
		result, err := userCollection.UpdateOne(ctx, bson.M{"id": objectId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			})
			return
		}

		var updatedUser model.User

		if result.MatchedCount == 1 {
			err := userCollection.FindOne(ctx, bson.M{"id": objectId}).Decode(&updatedUser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    map[string]interface{}{"data": err.Error()},
				})
				return
			}
		}

		c.JSON(http.StatusOK, responses.UserResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": updatedUser},
		})
	}
}
