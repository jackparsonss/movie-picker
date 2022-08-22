package cmd

import (
	"fmt"
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
		fmt.Println("http://localhost:8080")

		// Serve frontend static files
		r.Use(static.Serve("/", static.LocalFile("./client/build", true)))
		r.GET("/api/list", api.List)
		log.Fatal(r.Run())
	},
}

func init() {
	RootCmd.AddCommand(WebCmd)
}
