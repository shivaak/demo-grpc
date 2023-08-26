package entity

import (
	"fmt"
	"github.com/google/uuid"
)

type Book struct {
	Id    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	Count int32     `json:"count"`
}

func (b *Book) String() string {
	return fmt.Sprintf("Id: #{b.Id} Title: #{b.Title} Count: #{b.Count}")
}
