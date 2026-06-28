package users

import (
	"api/internal/db"
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const COLLECTION_NAME = "users"

type User struct {
	Name     string        `bson:"name"`
	Dob      string        `bson:"dob"`
	Username string        `bson:"username"`
	ID       bson.ObjectID `bson:"_id,omitempty"`
}

func (user *User) CreateUser(ctx context.Context) (*User, error) {
	if user.ID.IsZero() {
		user.ID = bson.NewObjectID()
	}
	coll := db.GetDB().GetCollection(ctx, COLLECTION_NAME)
	_, err := coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (user *User) UpdateUser(ctx context.Context) (*User, error) {
	if user.ID.IsZero() {
		return nil, nil
	}
	coll := db.GetDB().GetCollection(ctx, COLLECTION_NAME)
	result := coll.FindOneAndUpdate(ctx,
		bson.M{"_id": user.ID},
		bson.M{"$set": user},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)
	if err := result.Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}
