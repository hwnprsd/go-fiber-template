package users

import (
	"fmt"

	"flaq.club/backend/configs"
	"flaq.club/backend/models"
	"flaq.club/backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, models.USERS)

// Get user profile
func GetProfile(data utils.RequestBody) interface{} {
	return "User Profile Ideally"
}

func UpdateUser(data utils.RequestBody) interface{} {
	return "User updated successfully"
}

func CreateUser(data models.User) interface{} {
	ctx, cancel := utils.GetContext()
	defer cancel()
	user := models.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Username:  data.Username,
		Id:        primitive.NewObjectID(),
	}

	result, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		fmt.Println(err)
		panic("Error creating user - User Service")
	}
	return result
}

// Find a user by given ID
func FindUserById(id string) models.User {
	ctx, cancel := utils.GetContext()
	defer cancel()

	var user models.User

	err := userCollection.FindOne(ctx, bson.M{"id": utils.ObjId(id)}).Decode(&user)
	if err != nil {
		fmt.Println(id, err)
		panic("Cannot find user")
	}
	return user
}

func GetOneUser(data utils.RequestBody) interface{} {
	body := data.Data.(*OneUserDto)
	return FindUserById(body.Id)
}

func GetAllUsers(data utils.RequestBody) interface{} {
	ctx, cancel := utils.GetContext()
	defer cancel()
	result, err := userCollection.Find(ctx, bson.M{})

	if err != nil {
		fmt.Println(err)
		panic("Cannot find users - User Service ")
	}

	var users []models.User

	defer result.Close(ctx)
	for result.Next(ctx) {
		var user models.User
		if err := result.Decode(&user); err != nil {
			continue
		}
		users = append(users, user)
	}

	return users
}
