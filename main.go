package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/okumawat/go-concurrency/cache"
	"github.com/okumawat/go-concurrency/dao"
	"github.com/okumawat/go-concurrency/model"
)

var random = rand.New(rand.NewSource(time.Now().UnixMilli()))

func CacheDemo() {
	wg := &sync.WaitGroup{}
	mutex := &sync.RWMutex{}

	//Creating channel for cache and db output
	cacheChannel := make(chan model.Book)
	dbChannel := make(chan model.Book)

	for i := 0; i < 15; i++ {
		bookId := random.Intn(10)
		//Define no of tasks in waitgroup
		wg.Add(2)

		//Get data from cache
		go func(id int, wg *sync.WaitGroup, mt *sync.RWMutex, cacheChannel chan<- model.Book) {
			if b, ok := cache.GetBookById(id, mt); ok {
				cacheChannel <- b
			} else {
				fmt.Printf("Id %v Not found in cache!!\n", id)
			}
			wg.Done()
		}(bookId, wg, mutex, cacheChannel)

		//Get data from db
		go func(id int, wg *sync.WaitGroup, mt *sync.RWMutex, dbChannel chan<- model.Book) {
			if b, ok := dao.GetBookById(id, mt); ok {
				dbChannel <- b
			} else {
				fmt.Printf("Id %v Not found in db!!\n", id)
			}
			wg.Done()
		}(bookId, wg, mutex, dbChannel)

		//Print book data
		go func(dbChannel chan model.Book, cacheChannel chan model.Book) {

			select {
			case b := <-cacheChannel:
				fmt.Println("From cache:", b.String())
				<-dbChannel

			case b := <-dbChannel:
				fmt.Println("From databse:", b.String())
			}

		}(dbChannel, cacheChannel)

		time.Sleep(110 * time.Millisecond)

	}

	//Wait for tasks to finish
	wg.Wait()
	fmt.Println("Processing done.")
}

func main() {
	CacheDemo()
	//channel.ChannelDemo()
}
