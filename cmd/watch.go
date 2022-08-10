package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"movie/db"
	"strconv"
)

var WatchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Moves move from movie list to watched list",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])

		if err != nil {
			log.Fatalln("Failed to parse: ", id)
		}

		// get all movies from movie list
		movies, err := db.AllMovies(db.MoviesBucket)

		if err != nil {
			log.Fatalln(err)
		}

		if id <= 0 || id > len(movies) {
			log.Fatalln("Invalid movie number:", id)
		}
		movie := movies[id-1]

		// delete movie from movie list
		err = db.DeleteMovie(movie.Key, db.MoviesBucket)
		if err != nil {
			log.Fatalln(err)
		}

		// add to watched list bucket
		_, err = db.CreateMovie(movie.Value, db.WatchedBucket)

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("You have watched \"%s\"!\n", movie.Value)
	},
}

func init() {
	RootCmd.AddCommand(WatchCmd)
}
