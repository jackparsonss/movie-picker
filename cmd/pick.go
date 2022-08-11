package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jackparsonss/movie/db"
	"github.com/spf13/cobra"
	"log"
	"math/rand"
	"strings"
)

var PickCmd = &cobra.Command{
	Use:   "pick",
	Short: "Randomly picks a movie",
	Run: func(cmd *cobra.Command, args []string) {
		movies, err := db.AllMovies(db.MoviesBucket)

		if err != nil {
			log.Fatalln(err)
		}

		if len(movies) == 0 {
			fmt.Println("You have no movies...")
			return
		}

		for {
			randomIndex := rand.Intn(len(movies))
			d := color.New(color.FgBlue, color.Bold, color.Underline)
			fmt.Printf("Would you like to watch ")
			d.Printf("\"%s\"?", movies[randomIndex].Value)
			fmt.Printf("(yes/no): ")

			var ans string
			fmt.Scanln(&ans)
			if strings.ToLower(ans) == "yes" || strings.ToLower(ans) == "y" {
				title := movies[randomIndex].Value

				// delete from current bucket
				err := db.DeleteMovie(movies[randomIndex].Key, db.MoviesBucket)
				if err != nil {
					log.Fatalln(err)
				}

				// add to watched list bucket
				_, err = db.CreateMovie(title, db.WatchedBucket)

				if err != nil {
					log.Fatalln(err)
				}
				fmt.Println("Moved to watched list...")
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(PickCmd)
}
