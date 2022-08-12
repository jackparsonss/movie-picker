package main

import (
	"github.com/jackparsonss/movie/cmd"
	"github.com/jackparsonss/movie/db"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	err := db.Connect("/Users/jackparsons/Storage/bolt/movies.db")
	must(err)

	err = cmd.RootCmd.Execute()
	must(err)
}

func must(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
