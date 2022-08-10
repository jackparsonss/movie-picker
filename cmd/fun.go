package cmd

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"log"
	"movie/db"
)

var FunCmd = &cobra.Command{
	Use:   "fun",
	Short: "interactive cli",
	Run: func(cmd *cobra.Command, args []string) {
		menuOptions := []string{"• Movie List", "• Watched Movie List", "• Add a movie", "• Add a watched movie", "• EXIT"}
		prompt := promptui.Select{
			Label: "Select Your Option",
			Items: menuOptions,
		}

		for {
			_, result, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			switch result {
			case menuOptions[0]:
				handleList()
			case menuOptions[1]:
				handleWatchedList()
			case menuOptions[2]:
				handleAdd()
			case menuOptions[3]:
				handleAddWatched()
			default:
				return
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(FunCmd)
}

func handleList() {
	for {
		movies, err := db.AllMovies(db.MoviesBucket)

		if err != nil {
			log.Fatalln(err)
		}

		movieTitles := []string{"• GO BACK"}
		for _, movie := range movies {
			movieTitles = append(movieTitles, "• "+movie.Value)
		}
		prompt := promptui.Select{
			Label: "Your movies",
			Items: movieTitles,
		}

		i, result, err := prompt.Run()

		if i == 0 {
			break
		}

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		options := []string{"• Watch Movie", "• Delete Movie", "• Go Back"}

		prompt = promptui.Select{
			Label: fmt.Sprintf("What would you like to do to %s", result),
			Items: options,
		}

		newI, _, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		movie := movies[i-1]
		if newI == 0 {
			err := db.DeleteMovie(movie.Key, db.MoviesBucket)
			if err != nil {
				log.Fatalln(err)
			}
			// add to watched list bucket
			_, err = db.CreateMovie(movie.Value, db.WatchedBucket)

			if err != nil {
				log.Fatalln(err)
			}
		} else if newI == 1 {
			err := db.DeleteMovie(movie.Key, db.MoviesBucket)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func handleWatchedList() {
	for {
		movies, err := db.AllMovies(db.WatchedBucket)

		if err != nil {
			log.Fatalln(err)
		}

		movieTitles := []string{"• GO BACK"}
		for _, movie := range movies {
			movieTitles = append(movieTitles, "• "+movie.Value)
		}
		prompt := promptui.Select{
			Label: "Your movies",
			Items: movieTitles,
		}

		i, result, err := prompt.Run()

		if i == 0 {
			break
		}

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		options := []string{"• Unwatch Movie", "• Delete Movie", "• Go Back"}

		prompt = promptui.Select{
			Label: fmt.Sprintf("What would you like to do to %s", result),
			Items: options,
		}

		newI, _, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		movie := movies[i-1]
		if newI == 0 {
			err := db.DeleteMovie(movie.Key, db.WatchedBucket)
			if err != nil {
				log.Fatalln(err)
			}
			// add to watched list bucket
			_, err = db.CreateMovie(movie.Value, db.MoviesBucket)

			if err != nil {
				log.Fatalln(err)
			}
		} else if newI == 1 {
			err := db.DeleteMovie(movie.Key, db.WatchedBucket)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
}

func handleAdd() {
	fmt.Println("Listing movies...")
}

func handleAddWatched() {
	fmt.Println("Listing movies...")
}
