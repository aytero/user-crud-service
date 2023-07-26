package model

import (
	"encoding/json"
	"strconv"
)

// todo json to dto, bson in model

type StringInt int

// UserInfo is a public info about user, without the password hash
type UserInfo struct {
	InsertedId string    `json:"-" bson:"_id,omitempty"`
	Id         string    `json:"id" bson:"id,omitempty" binding:"required"`
	IsActive   bool      `json:"isActive" bson:"is_active,omitempty"` // binding:"required"`
	Balance    string    `json:"balance" bson:"balance,omitempty"`
	Age        StringInt `json:"age" bson:"age,omitempty" binding:"required"`
	Name       string    `json:"name" bson:"name,omitempty" binding:"required"`
	Gender     string    `json:"gender" bson:"gender,omitempty" binding:"required"`
	Company    string    `json:"company" bson:"company,omitempty" binding:"required"`
	Email      string    `json:"email" bson:"email,omitempty" binding:"required"`
	Phone      string    `json:"phone" bson:"phone,omitempty" binding:"required"`
	Address    string    `json:"address" bson:"address,omitempty" binding:"required"`
	About      string    `json:"about" bson:"about,omitempty" binding:"required"`
	Registered string    `json:"registered" bson:"registered,omitempty" binding:"required"`
	Latitude   float64   `json:"latitude" bson:"latitude,omitempty"`
	Longitude  float64   `json:"longitude" bson:"longitude,omitempty"`
	Tags       []string  `json:"tags" bson:"tags,omitempty"`
	Friends    []Friends `json:"friends" bson:"friend,omitempty"`
	Data       string    `json:"data" bson:"data,omitempty" binding:"required"`
}

type User struct {
	InsertedId string    `json:"-" bson:"_id,omitempty"`
	Id         string    `json:"id" bson:"id,omitempty" binding:"required"`
	Password   string    `json:"password" bson:"password,omitempty" binding:"required"`
	IsActive   bool      `json:"isActive" bson:"is_active,omitempty"`
	Balance    string    `json:"balance" bson:"balance,omitempty"`
	Age        StringInt `json:"age" bson:"age,omitempty" binding:"required"`
	Name       string    `json:"name" bson:"name,omitempty" binding:"required"`
	Gender     string    `json:"gender" bson:"gender,omitempty" binding:"required"`
	Company    string    `json:"company" bson:"company,omitempty" binding:"required"`
	Email      string    `json:"email" bson:"email,omitempty" binding:"required"`
	Phone      string    `json:"phone" bson:"phone,omitempty" binding:"required"`
	Address    string    `json:"address" bson:"address,omitempty" binding:"required"`
	About      string    `json:"about" bson:"about,omitempty" binding:"required"`
	Registered string    `json:"registered" bson:"registered,omitempty" binding:"required"`
	Latitude   float64   `json:"latitude" bson:"latitude,omitempty"`
	Longitude  float64   `json:"longitude" bson:"longitude,omitempty"`
	Tags       []string  `json:"tags" bson:"tags,omitempty"`
	Friends    []Friends `json:"friends" bson:"friend,omitempty"`
	Data       string    `json:"data" bson:"data,omitempty" binding:"required"`
}

type Friends struct {
	Id   int    `json:"id" bson:"id,omitempty"`
	Name string `json:"name" bson:"name,omitempty"`
}

type UpdateUser struct {
	InsertedId string    `json:"-" bson:"_id,omitempty"`
	Id         string    `json:"id,omitempty" bson:"id,omitempty"`
	Password   string    `json:"password,omitempty" bson:"password,omitempty"`
	IsActive   bool      `json:"isActive,omitempty" bson:"is_active,omitempty"`
	Balance    string    `json:"balance,omitempty" bson:"balance,omitempty"`
	Age        StringInt `json:"age,omitempty" bson:"age,omitempty"`
	Name       string    `json:"name,omitempty" bson:"name,omitempty"`
	Gender     string    `json:"gender,omitempty" bson:"gender,omitempty"`
	Company    string    `json:"company,omitempty" bson:"company,omitempty"`
	Email      string    `json:"email,omitempty" bson:"email,omitempty"`
	Phone      string    `json:"phone,omitempty" bson:"phone,omitempty"`
	Address    string    `json:"address,omitempty" bson:"address,omitempty"`
	About      string    `json:"about,omitempty" bson:"about,omitempty"`
	Registered string    `json:"registered,omitempty" bson:"registered,omitempty"`
	Latitude   float64   `json:"latitude,omitempty" bson:"latitude,omitempty"`
	Longitude  float64   `json:"longitude,omitempty" bson:"longitude,omitempty"`
	Tags       []string  `json:"tags,omitempty" bson:"tags,omitempty"`
	Friends    []Friends `json:"friends,omitempty" bson:"friend,omitempty"`
	Data       string    `json:"data,omitempty" bson:"data,omitempty"`
}

type UpdateFriends struct {
	Id   int    `json:"id,omitempty" bson:"id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}

func (st *StringInt) UnmarshalJSON(b []byte) error {
	var item interface{}
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}
	switch v := item.(type) {
	case int:
		*st = StringInt(v)
	case float64:
		*st = StringInt(int(v))
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*st = StringInt(i)
	}
	return nil
}
