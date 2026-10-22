package users

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"api/internal/db"
)

type User struct {
	Name     string             `json:"name" bson:"name"`
	Dob      string             `json:"dob" bson:"dob"`
	Username string             `json:"username" bson:"username"`
	Id       primitive.ObjectID `json:"id" bson:"id"`
}

func getUserCollection(mongoDB *mongo.Database) *mongo.Collection {
	return mongoDB.Collection("users")
}

func (u *User) GetByID(mongoDB *mongo.Database, ctx context.Context, id primitive.ObjectID) (*User, error) {
	var user User

	if err := getUserCollection(mongoDB).FindOne(ctx, bson.M{"_id": id}).Decode(&user); errors.Is(err, mongo.ErrNoDocuments) {
		return nil, db.DBError(db.NO_DOCUMENTS_ERROR, "User not found", bson.M{"id": id})
	}

	return &user, nil
}
