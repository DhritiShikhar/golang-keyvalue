package main

import (
	"errors"
	"log"
	"time"

	"github.com/coreos/bbolt"
)

// Kosh is the key value store
type Kosh struct {
	db *bolt.DB
}

// open connection
func Open() (*Kosh, error) {
	opts := &bolt.Options{
		Timeout: 1 * time.Second,
	}

	// Open file which is present in current directory
	// If it doesnot exist then it will be created
	db, err := bolt.Open("kosh.db", 0600, opts)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &Kosh{db: db}, nil
}

// close connection
func (k *Kosh) Close() error {
	return k.db.Close()
}

// Add data
func (k *Kosh) Add(key string, value interface{}) error {
	if key == "" {
		return errors.New("key not found")
	}
	if value == nil {
		return errors.New("value not found")
	}

	err = k.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(world)

	})

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
