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

func StoreThings(s Storer[string, int]) error {
	return s.Put("foo", 2)
}

func NewKVStore[K comparable, V any]() *KVStore[K, V] {
	return &KVStore[K, V]{
		data: make(map[K]V),
	}
}

type Block struct{}
type Transaction struct{}

func main() {
	_ = NewKVStore[string, *Block]()
	_ = NewKVStore[string, *Transaction]()
}
