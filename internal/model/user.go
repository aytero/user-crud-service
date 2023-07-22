package model

import "time"

type User struct {
    Id         string     `json:"id"`
    Password   string     `json:"password"`
    IsActive   bool       `json:"isActive"`
    Balance    string     `json:"balance"`
    Age        int32      `json:"age,string"`
    Name       string     `json:"name"`
    Gender     string     `json:"gender"`
    Company    string     `json:"company"`
    Email      string     `json:"email"`
    Phone      string     `json:"phone"`
    Address    string     `json:"address"`
    About      string     `json:"about"`
    Registered *time.Time `json:"registered"`
    //CompletedTime *time.Time `json:"completed_time,omitempty" db:"completed_time"`
    Latitude  float64   `json:"latitude"`
    Longitude float64   `json:"longitude"`
    Tags      []string  `json:"tags"`
    Friends   []Friends `json:"friends"`
    Data      string    `json:"data"`
}

type Friends struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
}
