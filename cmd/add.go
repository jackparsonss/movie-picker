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
		task := strings.Join(args, " ")
		_, err := db.CreateMovie(task)

		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Successfully added \"%s\" to your movie list\n!", task)
	},
}

func init() {
	RootCmd.AddCommand(AddCmd)
}
