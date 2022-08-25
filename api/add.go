package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackparsonss/movie/db"
)

type movie struct {
	Title string `json:"title"`
}

type movieWid struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}

func Add(c *gin.Context) {
	var newMovie movie

	if err := c.BindJSON(&newMovie); err != nil {
		c.Error(err)
		return
	}
	// add to watched list bucket
	key, err := db.CreateMovie(newMovie.Title, db.MoviesBucket)

	if err != nil {
		c.Error(err)
		return
	}

	ret := movieWid{
		Key:   key,
		Value: newMovie.Title,
	}

	c.JSON(200, ret)
}
