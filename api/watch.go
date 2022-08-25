package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackparsonss/movie/db"
	"strconv"
)

func Watch(c *gin.Context) {
	movies, err := db.AllMovies(db.MoviesBucket)

	if err != nil {
		c.Error(err)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.Error(err)
		return
	}

	title := movies[id].Value
	ret := movies[id]

	// delete from current bucket
	err = db.DeleteMovie(movies[id].Key, db.MoviesBucket)
	if err != nil {
		c.Error(err)
		return
	}

	// add to watched list bucket
	_, err = db.CreateMovie(title, db.WatchedBucket)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, ret)
}
