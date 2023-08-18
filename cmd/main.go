package main

import (
	"net/http"
	"os"
	"time"

	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func main() {
	// var logger = zerolog.New(os.Stdout).With().Timestamp().Logger().Level(zerolog.InfoLevel)
	// if os.Getenv("DEBUG") != "" {
	var logger = zerolog.New(os.Stdout).With().Timestamp().Logger().Level(zerolog.DebugLevel)
	// }
	var Port string
	lastRunTime := time.Now().Format("2006-01-02 15:04:05")
	if os.Getenv("PORT") != "" {
		Port = os.Getenv("PORT")
	} else {
		Port = "8080"
	}
	logger.Info().Msg("Starting...")
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.SetTrustedProxies([]string{"::1"})
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	r.Use(ginzerolog.Logger("gin"))
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"last_run": lastRunTime,
			"route":    "healthz",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"last_run": lastRunTime,
			"route":    "home",
		})
	})
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"last_run": lastRunTime,
			"route":    "test",
		})
	})
	logger.Info().Msg("Listening on port " + Port)
	r.Run(":" + Port)
}
