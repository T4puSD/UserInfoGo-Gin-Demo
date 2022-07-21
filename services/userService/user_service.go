package userService

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"user-info-service/db"
	"user-info-service/model"
	"user-info-service/services/authservice"
	"user-info-service/utils"
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

func RegisterNewUser(user *model.User) (*mongo.InsertOneResult, error) {
	if err := validateUserModel(user); err != nil {
		return nil, err
	}

	if emailExist, err := isExistingEmail(&user.Email); err != nil && emailExist {
		return nil, errors.New("Email already exists!")
	}

	hashedPassword, err := authservice.GetHashedPassword(user.Password)
	if utils.IsBlank(hashedPassword) && err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	newUser := model.User{
		Id:       primitive.NewObjectID(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	ctx, cancel := getContext()
	defer cancel()
	return collection.InsertOne(ctx, newUser)
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

func GetByEmail(email *string) (*model.User, error) {
	ctx, cancel := getContext()
	defer cancel()

	var user model.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found with provided email")
	}

	return &user, nil
}

func isExistingEmail(email *string) (bool, error) {
	ctx, cancel := getContext()
	defer cancel()

	count, err := collection.CountDocuments(ctx, bson.M{"email": email})
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
