package db

import (
	"encoding/binary"
	"fmt"
	"github.com/boltdb/bolt"
	"time"
)

var MoviesBucket = []byte("movies")
var WatchedBucket = []byte("watched")
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

	// create movies bucket
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(MoviesBucket)

		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		return err
	})

	if err != nil {
		return err
	}

	// create watched bucket
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(WatchedBucket)

		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		return err
	})

	if err != nil {
		return err
	}

	return nil
}

func CreateMovie(task string, bucket []byte) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
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

func AllMovies(bucket []byte) ([]Movie, error) {
	var movies []Movie
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			movie := Movie{
				Key:   btoi(k),
				Value: string(v),
			}
			movies = append(movies, movie)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return movies, nil
}

func DeleteMovie(key int, bucket []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
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
