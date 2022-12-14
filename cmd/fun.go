package cmd

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/jackparsonss/movie/db"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var FunCmd = &cobra.Command{
	Use:   "fun",
	Short: "interactive cli",
	Run: func(cmd *cobra.Command, args []string) {
		menuOptions := []string{"• Pick Movie", "• Movie List", "• Watched Movie List", "• Add a movie", "• Add a watched movie", "✕ Exit"}
		prompt := promptui.Select{
			Label: "Select Your Option",
			Items: menuOptions,
			Size:  6,
		}

		for {
			_, result, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			switch result {
			case menuOptions[0]:
				pickMovie()
			case menuOptions[1]:
				handleList()
			case menuOptions[2]:
				handleWatchedList()
			case menuOptions[3]:
				handleAdd()
			case menuOptions[4]:
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

func pickMovie() {
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

		if strings.ToLower(result) == "yes" || strings.ToLower(result) == "y" {
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
}

func handleList() {
	for {
		movies, err := db.AllMovies(db.MoviesBucket)

		if err != nil {
			log.Fatalln(err)
		}

		movieTitles := []string{"✕ Cancel"}
		for _, movie := range movies {
			movieTitles = append(movieTitles, "• "+movie.Value)
		}
		prompt := promptui.Select{
			Label: "Your movies:",
			Items: movieTitles,
			Size:  30,
		}

		i, _, err := prompt.Run()

		if i == 0 {
			break
		}

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		options := []string{"• Watch Movie", "• Delete Movie", "✕ Cancel"}

		prompt = promptui.Select{
			Label: fmt.Sprintf("What would you like to do to %s", movies[i-1].Value),
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

		movieTitles := []string{"✕ Cancel"}
		for _, movie := range movies {
			movieTitles = append(movieTitles, "• "+movie.Value)
		}
		prompt := promptui.Select{
			Label: "Your watched movies:",
			Items: movieTitles,
			Size:  30,
		}

		i, _, err := prompt.Run()

		if i == 0 {
			break
		}

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		options := []string{"• Unwatch Movie", "• Delete Movie", "✕ Cancel"}

		prompt = promptui.Select{
			Label: fmt.Sprintf("What would you like to do to %s", movies[i-1].Value),
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
	handleAddMovie(db.MoviesBucket)
}

func handleAddWatched() {
	handleAddMovie(db.WatchedBucket)
}

func handleAddMovie(bucket []byte) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the movie title: ")
	movieTitle, _ := reader.ReadString('\n')
	movieTitle = movieTitle[:len(movieTitle)-1]

	_, err := db.CreateMovie(movieTitle, bucket)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Succesfully added \"%s\"\n", movieTitle)
}
