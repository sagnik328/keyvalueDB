package main

import "sync"

type Storer[K comparable, V any] interface {
	Put(K, V) error
	Get(K) (V, error)
	Update(K, V) error
	Delete(K) (V, error)
}

type KVStore[K comparable, V any] struct {
	mu   sync.Mutex
	data map[K]V
}

func StoreThings(s Storer) error {
	return s.Put("foo", []byte("Br"))
}

func NewKVStore() *KVStore[string, int] {
	return &KVStore[string, int]{
		data: make(map[string]int),
	}
}

func main() {

}
