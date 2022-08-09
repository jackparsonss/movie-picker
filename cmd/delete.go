package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"movie/db"
	"strconv"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a movie",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, v := range args {
			id, err := strconv.Atoi(v)

			if err != nil {
				log.Fatalln("Failed to parse: ", v)
			}

			ids = append(ids, id)
		}

		movies, err := db.AllMovies()

		if err != nil {
			log.Fatalln(err)
		}

		for _, id := range ids {
			if id <= 0 || id > len(movies) {
				fmt.Println("Invalid movie number:", id)
				continue
			}
			movie := movies[id-1]

			err := db.DeleteMovie(movie.Key)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("Deleting movie \"%s\"\n", movie.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(DeleteCmd)
}
