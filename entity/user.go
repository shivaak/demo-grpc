package entity

import (
	"fmt"
	"github.com/google/uuid"
)

type User struct {
	Id      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Age     int32     `json:"age"`
	City    string    `json:"city"`
	Country string    `json:"country"`
}

func (u *User) String() string {
	return fmt.Sprintf("Id: #{u.Id}, Name: #{u.Name}")
}
