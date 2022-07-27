package cache

import (
	"sync"

	"github.com/okumawat/go-concurrency/model"
)

var cache = map[int]model.Book{}

func GetBookById(id int, mt *sync.RWMutex) (model.Book, bool) {
	mt.RLock()
	b, ok := cache[id]
	mt.RUnlock()
	return b, ok
}

func UpdateCache(id int, b model.Book, mt *sync.RWMutex) {
	mt.Lock()
	//defer mt.Unlock()
	cache[id] = b
	mt.Unlock()
}
