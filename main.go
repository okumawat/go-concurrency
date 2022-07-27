package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/okumawat/go-concurrency/cache"
	"github.com/okumawat/go-concurrency/dao"
)

var random = rand.New(rand.NewSource(time.Now().UnixMilli()))

func main() {
	wg := &sync.WaitGroup{}
	mutex := &sync.RWMutex{}

	for i := 0; i < 15; i++ {
		bookId := random.Intn(10)
		//Define no of tasks in waitgroup
		wg.Add(2)

		//Get data from cache
		go func(id int, wg *sync.WaitGroup, mt *sync.RWMutex) {
			if b, ok := cache.GetBookById(id, mt); ok {
				fmt.Println("From cache:", b.String())
			} else {
				fmt.Printf("Id %v Not found in cache!!\n", id)
			}
			wg.Done()
		}(bookId, wg, mutex)

		//Get data from db
		go func(id int, wg *sync.WaitGroup, mt *sync.RWMutex) {
			if b, ok := dao.GetBookById(id, mt); ok {
				fmt.Println("From database:", b.String())
			} else {
				fmt.Printf("Id %v Not found in db!!\n", id)
			}
			wg.Done()
		}(bookId, wg, mutex)

		time.Sleep(110 * time.Millisecond)

	}

	//Wait for tasks to finish
	wg.Wait()
	fmt.Println("Processing done.")
}
