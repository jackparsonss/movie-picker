package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jackparsonss/movie/db"
	"github.com/manifoldco/promptui"
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
			d := color.New(color.FgBlue, color.Bold, color.Underline).SprintFunc()
			menuOptions := []string{"Yes", "No"}
			prompt := promptui.Select{
				Label: "Would you like to watch " + d(movies[randomIndex].Value),
				Items: menuOptions,
			}

			_, result, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			if strings.ToLower(result) == "yes" {
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
				return
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(PickCmd)
}
