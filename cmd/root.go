package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "movie",
	Short: "Movie is a CLI task manager",
}

func init() {
	RootCmd.PersistentFlags().Bool("watched", false, "A whether to use movie list or watched list")
}
