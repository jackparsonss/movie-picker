package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackparsonss/movie/db"
	"strconv"
)

func Move(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}

	bucket := c.Query("bucket")
	var newMovie movie

	if err = c.BindJSON(&newMovie); err != nil {
		c.Error(err)
		return
	}

	if bucket == "watched" {
		err = db.DeleteMovie(id, db.WatchedBucket)
		if err != nil {
			c.Error(err)
			return
		}
		_, err = db.CreateMovie(newMovie.Title, db.MoviesBucket)

		if err != nil {
			c.Error(err)
			return
		}
	} else if bucket == "toWatch" {
		err = db.DeleteMovie(id, db.MoviesBucket)
		if err != nil {
			c.Error(err)
			return
		}
		_, err = db.CreateMovie(newMovie.Title, db.WatchedBucket)

		if err != nil {
			c.Error(err)
			return
		}
	} else {
		if err != nil {
			c.Error(errors.New("need to pass in 'bucket' parameter"))
			return
		}
	}

	c.JSON(200, ret{
		Id:     id,
		Bucket: bucket,
	})
}
