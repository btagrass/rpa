package mgt

import (
	"github.com/btagrass/gobiz/cmw"
	s "github.com/btagrass/gobiz/svc"
	"github.com/btagrass/gobiz/sys/mgt"
	"github.com/btagrass/gobiz/sys/svc"
	"github.com/gin-gonic/gin"
)

func Mgt() *gin.Engine {
	e := mgt.Mgt()
	m := e.Group("/mgt").Use(cmw.Auth(s.Use[*svc.UserSvc]().Perm, s.Use[*svc.UserSvc]().SignedKey), cmw.Visit(s.Use[*svc.VisitSvc]()))
	{
		m.GET("/processes", ListProcesses)
		m.GET("/processes/:id", GetProcess)
		m.DELETE("/processes/:ids", RemoveProcesses)
		m.POST("/processes", SaveProcess)
		m.GET("/processes/:id/steps", ListSteps)
		m.DELETE("/processes/steps/:ids", RemoveSteps)
		m.POST("/processes/:id/steps", SaveStep)
	}
	return e
}
