package main

import (
	"log"
	"math/rand"
	"movie/cmd"
	"movie/db"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	err := db.Connect("/Users/jackparsons/Storage/bolt/tasks.db")
	must(err)

	err = cmd.RootCmd.Execute()
	must(err)
}

func must(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
