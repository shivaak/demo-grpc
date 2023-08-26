package model

import "github.com/google/uuid"

// User is the repository layer model for a user
type User struct {
	Id      uuid.UUID
	Name    string
	Age     int32
	City    string
	Country string
}

type Book struct {
	Id    uuid.UUID
	Title string
	Count int32
}
