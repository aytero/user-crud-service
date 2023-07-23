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

	Id         string    `json:"id" bson:"id"`
	Password   string    `json:"password" bson:"password"`
	IsActive   bool      `json:"isActive" bson:"is_active"`
	Balance    string    `json:"balance" bson:"balance"`
	Age        StringInt `json:"age" bson:"age"`
	Name       string    `json:"name" bson:"name"`
	Gender     string    `json:"gender" bson:"gender"`
	Company    string    `json:"company" bson:"company"`
	Email      string    `json:"email" bson:"email"`
	Phone      string    `json:"phone" bson:"phone"`
	Address    string    `json:"address" bson:"address"`
	About      string    `json:"about" bson:"about"`
	Registered string    `json:"registered" bson:"registered"`
	//Registered *time.Time `json:"registered,string" bson:"registered"`
	Latitude  float64   `json:"latitude" bson:"latitude"`
	Longitude float64   `json:"longitude" bson:"longitude"`
	Tags      []string  `json:"tags" bson:"tags"`
	Friends   []Friends `json:"friends" bson:"friend"`
	Data      string    `json:"data" bson:"data"`
}

type Friends struct {
	Id   int    `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
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
