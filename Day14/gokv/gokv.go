package gokv

import (
	"bytes"
	"encoding/gob"
	"errors"
	"github.com/boltdb/bolt"
)

var (
	ErrKeyNotFound = errors.New("Key not found")
	bucketName     = []byte("GoKV")
)

type KVStore struct {
	db *bolt.DB
}

func Open(path string) (*KVStore, error) {
	db, err := bolt.Open(path, 0640, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		return err
	})

	if err != nil {
		return nil, err
	} else {
		store := &KVStore{db}
		return store, nil
	}
}

func (store *KVStore) Put(key string, value interface{}) error {
	err := store.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(bucketName).Put([]byte(key), encode(value))
	})
	return err
}

func (store *KVStore) Get(key string, e interface{}) error {
	err := store.db.View(func(tx *bolt.Tx) error {
		value := tx.Bucket(bucketName).Get([]byte(key))
		if value == nil {
			return ErrKeyNotFound
		}

		decode(value, e)
		return nil
	})
	return err
}

func (store *KVStore) Delete(key string) error {
	err := store.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(bucketName).Delete([]byte(key))
	})
	return err
}

func encode(e interface{}) []byte {
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	encoder.Encode(e)
	return buffer.Bytes()
}

func decode(value []byte, e interface{}) {
	reader := bytes.NewReader(value)
	decoder := gob.NewDecoder(reader)
	decoder.Decode(e)
}
