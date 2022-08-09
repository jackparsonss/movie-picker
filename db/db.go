package db

import (
	"encoding/binary"
	"fmt"
	"github.com/boltdb/bolt"
	"time"
)

var moviesBucket = []byte("movies")
var db *bolt.DB

type Movie struct {
	Key   int
	Value string
}

func Connect(path string) error {
	var err error
	db, err = bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})

	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(moviesBucket)

		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		return err
	})
}

func CreateMovie(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(moviesBucket)
		id64, _ := b.NextSequence()
		id = int(id64)

		key := itob(id)
		return b.Put(key, []byte(task))
	})

	if err != nil {
		return -1, err
	}

	return id, nil
}

func AllMovies() ([]Movie, error) {
	var tasks []Movie
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(moviesBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			task := Movie{
				Key:   btoi(k),
				Value: string(v),
			}
			tasks = append(tasks, task)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func DeleteMovie(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(moviesBucket)
		return b.Delete(itob(key))
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
