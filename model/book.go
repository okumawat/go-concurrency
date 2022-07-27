package model

import "fmt"

type Book struct {
	Id          int
	Title       string
	Author      string
	PublishYear int
}

func (b Book) String() string {
	return fmt.Sprintf("\nBook{\n\tId:%v\n\tTitle:%v\n\tAuthor:%v\n\tPublishYear:%v\n}", b.Id, b.Title, b.Author, b.PublishYear)
}
