package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"math/rand"
	"movie/db"
	"strings"
)

var PickCmd = &cobra.Command{
	Use:   "pick",
	Short: "Randomly pick a movie",
	Run: func(cmd *cobra.Command, args []string) {
		movies, err := db.AllMovies()

		if err != nil {
			log.Fatalln(err)
		}

		if len(movies) == 0 {
			fmt.Println("You have no movies...")
			return
		}

		for {
			randomIndex := rand.Intn(len(movies))
			d := color.New(color.FgCyan, color.Bold, color.Underline)
			fmt.Printf("Would you like to watch ")
			d.Printf("\"%s\"?", movies[randomIndex].Value)
			fmt.Printf("(yes/no): ")

			var ans string
			fmt.Scanln(&ans)
			if strings.ToLower(ans) == "yes" || strings.ToLower(ans) == "y" {
				// TODO: delete from current bucket, add to watched bucket
				break
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(PickCmd)
}
