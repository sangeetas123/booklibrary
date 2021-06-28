package util

import (
	"encoding/json"
	redis "gopkg.in/redis.v5"
	"log"
)

var client *redis.Client

func InitializeRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	SetBookList(BookCollection)
}

func SetBookList(booklist *Collection) {
	err := client.Set("booklist",  booklist, 0).Err()
	if err != nil {
		log.Println("Couldn't save the booklist, Err: ", err)
	}
}

func GetBookList() *Collection {
	val, err := client.Get("booklist").Result()
	if err != nil {
		panic(err)
	}

	var books Collection
	err = json.Unmarshal([]byte(val), &books)
	if err != nil {
		panic(err)
	}

	return &books
}

