package dto

import "time"

type CreateUserRequest struct {
	Users []CreateUserDto `json:"users"`
}

type CreateUserResponse struct {
	Users []UserDto `json:"users"`
}

type CreateUserDto struct {
	Id         string       `json:"id,omitempty"`
	Password   string       `json:"password"`
	IsActive   bool         `json:"isActive"`
	Balance    string       `json:"balance"`
	Age        int32        `json:"age,string"`
	Name       string       `json:"name"`
	Gender     string       `json:"gender"`
	Company    string       `json:"company"`
	Email      string       `json:"email"`
	Phone      string       `json:"phone"`
	Address    string       `json:"address"`
	About      string       `json:"about"`
	Registered *time.Time   `json:"registered"`
	Latitude   float64      `json:"latitude"`
	Longitude  float64      `json:"longitude"`
	Tags       []string     `json:"tags"`
	Friends    []FriendsDto `json:"friends"`
	Data       string       `json:"data"`
}
