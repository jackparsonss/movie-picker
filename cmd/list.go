package cmd

import (
	"fmt"
	"github.com/jackparsonss/movie/db"
	"github.com/spf13/cobra"
	"log"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all movies",
	Run: func(cmd *cobra.Command, args []string) {
		bucket := getWatchedFlag(cmd)
		movies, err := db.AllMovies(bucket)

		if err != nil {
			log.Fatalln(err)
		}

		if len(movies) == 0 {
			fmt.Println("You have no movies...")
			return
		}

		fmt.Println("You have the following movies:")
		for i, movie := range movies {
			fmt.Printf("%d. %s\n", i+1, movie.Value)
		}

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
	},
}

func init() {
	RootCmd.AddCommand(ListCmd)
}
