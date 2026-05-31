package users


type User struct {
	Name string `json:"name" bson:"name"`
	Dob string `json:"dob" bson:"dob"`
	Username string `json:"username" bson:"username"`
}

