package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"movie/db"
	"strings"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a movie",
	Run: func(cmd *cobra.Command, args []string) {
		bucket := getWatchedFlag(cmd)

		movie := strings.Join(args, " ")
		_, err := db.CreateMovie(movie, bucket)

		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Successfully added \"%s\" to your movie list!\n", movie)
	},
}

func init() {
	RootCmd.AddCommand(AddCmd)
}
