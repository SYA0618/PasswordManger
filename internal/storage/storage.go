package storage

import (
	"encoding/json"
	"errors"
	"passwordManger/internal/model"

	"go.etcd.io/bbolt"
)

var db *bbolt.DB

const dbFile = "passman.db"
const bucketName = "passwords"

func InitDB() error {
	var err error
	db, err = bbolt.Open(dbFile, 0600, nil)
	if err != nil {
		panic(err)
	}

	return db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		return err
	})
}

func SaveEntry(entry model.Entry) error {
	return db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return errors.New("bucket not found")
		}
		data, err := json.Marshal(entry)
		if err != nil {
			return err
		}
		return b.Put([]byte(entry.Service), data)
	})
}

func GetEntry(service string) (model.Entry, error) {
	var entry model.Entry
	err := db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return errors.New("bucket not found")
		}
		data := b.Get([]byte(service))
		if data == nil {
			return errors.New("entry not found")
		}
		return json.Unmarshal(data, &entry)
	})
	return entry, err
}

func Close() {
	if db != nil {
		db.Close()
	}
}
