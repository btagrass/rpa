package api

import (
	"github.com/btagrass/gobiz/app"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Api() *gin.Engine {
	e := gin.Default()
	e.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"*"},
		AllowHeaders:    []string{"*"},
	}))
	e.Static("/data", app.DataDir)
	return e
}
