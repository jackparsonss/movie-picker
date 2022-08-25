package cmd

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jackparsonss/movie/api"
	"github.com/spf13/cobra"
	"log"
)

var WebCmd = &cobra.Command{
	Use:   "web",
	Short: "runs a simple web server",
	Run: func(cmd *cobra.Command, args []string) {
		gin.SetMode(gin.ReleaseMode)
		r := gin.New()
		r.Use(cors.New(cors.Config{
			AllowOrigins: []string{"http://localhost:3000"},
			AllowMethods: []string{"PUT", "GET"},
		}))
		fmt.Println("http://localhost:8080")

		// Serve frontend static files
		r.Use(static.Serve("/", static.LocalFile("./client/build", true)))

		r.GET("/api/list", api.List)
		r.PUT("/api/watch/:id", api.Watch)

		log.Fatal(r.Run())
	},
}

func init() {
	RootCmd.AddCommand(WebCmd)
}
