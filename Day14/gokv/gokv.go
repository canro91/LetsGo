package gokv

import (
	"bytes"
	"encoding/gob"
	"errors"
	"github.com/boltdb/bolt"
)

var (
	ErrKeyNotFound = errors.New("Key not found")
)

type KVStore struct {
	db *bolt.DB
}

func Open(path string) (*KVStore, error) {
	db, err := bolt.Open(path, 0640, nil)
	if err != nil {
		return nil, err
	}

	store := &KVStore{db}
	return store, nil
}

func (store *KVStore) Put(key string, value interface{}) error {
	err := store.db.Update(func(tx *bolt.Tx) error {
		bucketName := []byte("GoKV")
		bucket, _ := tx.CreateBucketIfNotExists(bucketName)

		err := bucket.Put(encode(key), encode(value))
		if err != nil {
			return err
		}

		return nil
	})
	return err
}

func (store *KVStore) Get(key string, e interface{}) error {
	err := store.db.View(func(tx *bolt.Tx) error {
		bucketName := []byte("GoKV")

		value := tx.Bucket(bucketName).Get(encode(key))
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
		bucketName := []byte("GoKV")

		err := tx.Bucket(bucketName).Delete(encode(key))
		if err != nil {
			return err
		}

		return nil
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
	buffer := bytes.Buffer{}
	buffer.Write(value)

	decoder := gob.NewDecoder(&buffer)
	decoder.Decode(e)
}
