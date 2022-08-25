package cmd

import (
	"log"
	"os/exec"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jackparsonss/movie/api"
	"github.com/spf13/cobra"
)

var WebCmd = &cobra.Command{
	Use:   "web",
	Short: "runs a simple web server",
	Run: func(cmd *cobra.Command, args []string) {
		gin.SetMode(gin.ReleaseMode)
		r := gin.New()

		config := cors.DefaultConfig()
		config.AllowOrigins = []string{"http://localhost:3000"}
		r.Use(cors.New(config))

		// MACOS ONLY
		exec.Command("open", "http://localhost:8080").Start()

		// Serve frontend static files
		r.Use(static.Serve("/", static.LocalFile("/Users/jackparsons/Storage/Repositories/movie-picker/client/build", true)))

		r.GET("/api/list", api.List)
		r.PUT("/api/watch/:id", api.Watch)
		r.PUT("/api/move/:id", api.Move)
		r.POST("/api/add", api.Add)
		r.DELETE("/api/delete/:id", api.Delete)

		log.Fatal(r.Run())
	},
}

func init() {
	RootCmd.AddCommand(WebCmd)
}
