package main

import (
	"fmt"
	"log"
	"sync"
)

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

func NewKVStore[K comparable, V any]() *KVStore[K, V] {
	//constructor for the KV Store
	return &KVStore[K, V]{
		data: make(map[K]V),
	}
}

func (s *KVStore[K, V]) Put(key K, value V) error {
	//implementation of the KVStorage
	//open a lock for reading and open a lock for writing.
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value

	return nil
}

func (s *KVStore[K, V]) Get(key K, value V) (V, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := s.data[key]
	if err {
		return value, fmt.Errorf("The key {%v} does not exists\n", key)
	}
	fmt.Printf("The key {%v} was found\n", data)
	return value, nil
}

func StoreThings(s Storer[string, int]) error {
	return s.Put("foo", 2)
}

//type Block struct{}
//type Transaction struct{}

func main() {
	//	_ = NewKVStore[string, *Block]()
	//	_ = NewKVStore[string, *Transaction]()

	store := NewKVStore[string, string]()
	if err := store.Put("foo", "bar"); err != nil {
		log.Fatal(err)
	}
}
