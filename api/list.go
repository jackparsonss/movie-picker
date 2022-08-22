package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackparsonss/movie/db"
)

func List(c *gin.Context) {
	watched := c.Query("type")
	t := db.MoviesBucket

	if watched == "watched" {
		t = db.WatchedBucket
	}
	movies, err := db.AllMovies(t)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, movies)
}
