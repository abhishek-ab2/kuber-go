package user

import (
	"api/internal/db/users"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserModel struct {
	Name     string        `json:"name"`
	Username string        `json:"username"`
	Dob      string        `json:"dob"`
	ID       bson.ObjectID `json:"id"`
}

func (model UserModel) toDB() users.User {
	return users.User{
		Name:     model.Name,
		Username: model.Username,
		Dob:      model.Dob,
		ID:       model.ID,
	}
}

type CreateUserModel struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Dob      string `json:"dob"`
}

type PatchUserModel struct {
	Name string `json:"name,omitempty"`
	Dob  string `json:"dob,omitempty"`
}
