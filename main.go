package main

import (
	"log"
	"movie/cmd"
	"movie/db"
)

func main() {
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
