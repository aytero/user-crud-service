package model

import (
	"encoding/json"
	"strconv"
)

// todo json to dto, bson in model

type StringInt int

type User struct {
	//InsertedId primitive.ObjectID `bson:"_id,omitempty"`
	InsertedId string `json:"-" bson:"_id,omitempty"`
	//Name     string             `bson:"name,omitempty"`
	//Email    string             `bson:"email,omitempty"`
	//Password string             `bson:"password,omitempty"`

	Id       string    `json:"id" bson:"id,omitempty"`
	Password string    `json:"password" bson:"password,omitempty"`
	IsActive bool      `json:"isActive" bson:"is_active,omitempty"`
	Balance  string    `json:"balance" bson:"balance,omitempty"`
	Age      StringInt `json:"age" bson:"age,omitempty"`
	Name     string    `json:"name" bson:"name,omitempty"`
	Gender   string    `json:"gender" bson:"gender,omitempty"`
	Company  string    `json:"company" bson:"company,omitempty"`
	Email    string    `json:"email" bson:"email,omitempty"`
	Phone    string    `json:"phone" bson:"phone,omitempty"`
	Address  string    `json:"address" bson:"address,omitempty"`
	About    string    `json:"about" bson:"about,omitempty"`

	Registered string `json:"registered" bson:"registered,omitempty"`
	//Registered *time.Time `json:"registered,string" bson:"registered"`
	Latitude  float64   `json:"latitude" bson:"latitude,omitempty"`
	Longitude float64   `json:"longitude" bson:"longitude,omitempty"`
	Tags      []string  `json:"tags" bson:"tags,omitempty"`
	Friends   []Friends `json:"friends" bson:"friend,omitempty"`
	Data      string    `json:"data" bson:"data,omitempty"`
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

// UnmarshalJSON create a custom unmarshal for the StringInt
// / this helps us check the type of our value before unmarshalling it
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
		///here convert the string into
		///an integer
		i, err := strconv.Atoi(v)
		if err != nil {
			///the string might not be of integer type
			///so return an error
			return err
		}
		*st = StringInt(i)
	}
	return nil
}
