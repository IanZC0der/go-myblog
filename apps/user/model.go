package user

import (
	"encoding/json"
	"time"
)

// User represents the user entity
type User struct {
	Id        int64 `json:"id"`
	CreatedAt int64 `json:"created_at"` //not datetime as it increases complexity (timezones, etc.)
	UpdatedAt int64 `json:"updated_at"`
	*CreateUserRequest
}

func NewUser(req *CreateUserRequest) *User {
	req.PasswordHash()
	return &User{
		CreatedAt:         time.Now().Unix(),
		CreateUserRequest: req,
	}
}

func (u *User) String() string {
	jsonUser, _ := json.Marshal(u)
	return string(jsonUser)
}

func (u *User) TableName() string {
	return "users"
}
