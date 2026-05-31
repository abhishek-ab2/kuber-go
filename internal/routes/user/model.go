package user

import "go.mongodb.org/mongo-driver/v2/bson"


type UserModel struct {
	Name string `bson:"name" json:"name"`
	Username string `bson:"name" json:"username"`
	Dob string `bson:"dob" json:"dob"`
	ID bson.ObjectID `bson:"_id" json:"id"`
}

type CreateUserModel struct {
	Name string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Dob string `json:"dob"`
}

type PatchUserModel struct {
	Name string `json:"name,omitempty"`
	Dob string `json:"dob,omitempty"`
}


