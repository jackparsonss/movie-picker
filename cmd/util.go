package cmd

import (
	"github.com/jackparsonss/movie/db"
	"github.com/spf13/cobra"
)

func getWatchedFlag(cmd *cobra.Command) []byte {
	watchedFlag, _ := cmd.Flags().GetBool("watched")
	bucket := db.MoviesBucket
	if watchedFlag {
		bucket = db.WatchedBucket
	}

	return bucket
}
