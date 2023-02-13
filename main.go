package main

import "sync"

type Storer interface {
	Put(string, []byte) error
	Get(string) ([]byte, error)
	Update(string, []byte) error
	Delete(string) (byte, error)
}

type KVStore struct {
	mu sync.Mutex
}

type StoreThings(s Storer) error {
	return s.Put("foo",[]byte("Br"))
}

func NewKVStore() *KVStore {
	return &KVStore{
		data: make(map[string][]byte)
	}
}

func main(){

}