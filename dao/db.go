package dao

import (
	"sync"
	"time"

	"github.com/okumawat/go-concurrency/cache"
	"github.com/okumawat/go-concurrency/model"
)

var books []model.Book = []model.Book{
	{Id: 1, Title: "Golang Concurrency", Author: "GoDev", PublishYear: 2022},
	{Id: 2, Title: "Java Concurrency", Author: "JavaDev", PublishYear: 2000},
	{Id: 3, Title: "C++ Concurrency", Author: "CPPDev", PublishYear: 2001},
	{Id: 4, Title: "Python Concurrency", Author: "PyDev", PublishYear: 2003},
	{Id: 5, Title: "C Concurrency", Author: "CDev", PublishYear: 2021},
	{Id: 6, Title: "Javascript Concurrency", Author: "JsDev", PublishYear: 2013},
}

func AddBook(b model.Book) {
	books = append(books, b)
}

func GetBooks() []model.Book {
	return books
}

func GetBookById(id int, mt *sync.Mutex) (model.Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.Id == id {
			cache.UpdateCache(id, b, mt)
			return b, true
		}
	}
	return model.Book{}, false
}
