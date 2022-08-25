package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackparsonss/movie/db"
	"strconv"
)

type ret struct {
	Id     int    `json:"key"`
	Bucket string `json:"bucket"`
}

func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	bucket := c.Query("bucket")

	if bucket == "watched" {
		err = db.DeleteMovie(id, db.WatchedBucket)
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
