package cmd

import (
	"github.com/spf13/cobra"
	"movie/db"
)

func getWatchedFlag(cmd *cobra.Command) []byte {
	watchedFlag, _ := cmd.Flags().GetBool("watched")
	bucket := db.MoviesBucket
	if watchedFlag {
		bucket = db.WatchedBucket
	}

	return bucket
}
