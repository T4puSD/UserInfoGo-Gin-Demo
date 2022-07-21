package userService

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"user-info-service/db"
	"user-info-service/model"
)

var validate = validator.New()
var collection = db.GetCollection("user")

func getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func validateUserModel(user *model.User) error {
	if err := validate.Struct(user); err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func Get(id string) (*model.User, error) {
	var user model.User

	ctx, cancel := getContext()
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := collection.FindOne(ctx, bson.M{"id": objectId}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func Save(user *model.User) error {
	if err := validateUserModel(user); err != nil {
		return err
	}

	newUser := model.User{
		Id:    primitive.NewObjectID(),
		Name:  user.Name,
		Email: user.Email,
	}

	ctx, cancel := getContext()
	defer cancel()

	_, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		return err
	}
	return nil
}

func Update(id string, user *model.User) error {
	if err := validateUserModel(user); err != nil {
		return err
	}

	ctx, cancel := getContext()
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{"name": user.Name, "email": user.Email}
	_, err = collection.UpdateOne(ctx, bson.M{"id": objectId}, bson.M{"$set": update})
	if err != nil {
		return err
	}
	return nil
}

func Delete(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	ctx, cancel := getContext()
	defer cancel()

	_, err = collection.DeleteOne(ctx, bson.M{"id": objectId})
	if err != nil {
		return err
	}
	return nil
}

func GetAll() (*[]model.User, error) {
	ctx, cancel := getContext()
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var users []model.User
	for cursor.Next(ctx) {
		var singleUser model.User
		if err := cursor.Decode(&singleUser); err != nil {
			return nil, err
		}
		users = append(users, singleUser)
	}

	// closing cursor
	err = cursor.Close(ctx)
	if err != nil {
		return nil, err
	}

	return &users, nil
}
