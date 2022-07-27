package cache

import (
	"sync"

	"github.com/okumawat/go-concurrency/model"
)

var cache = map[int]model.Book{}

func GetBookById(id int, mt *sync.Mutex) (model.Book, bool) {
	mt.Lock()
	b, ok := cache[id]
	mt.Unlock()
	return b, ok
}

func UpdateCache(id int, b model.Book, mt *sync.Mutex) {
	mt.Lock()
	//defer mt.Unlock()
	cache[id] = b
	mt.Unlock()
}
