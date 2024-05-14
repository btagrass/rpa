package mgt

import (
	"rpa/mdl"
	"rpa/svc"

	"github.com/btagrass/gobiz/r"
	s "github.com/btagrass/gobiz/svc"
	"github.com/btagrass/gobiz/utl"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func ListProcesses(c *gin.Context) {
	processes, count, err := s.Use[*svc.ProcessSvc]().List(r.Q(c))
	r.J(c, processes, count, err)
}

func GetProcess(c *gin.Context) {
	process, err := s.Use[*svc.ProcessSvc]().Get(c.Param("id"))
	r.J(c, process, err)
}

func RemoveProcesses(c *gin.Context) {
	err := s.Use[*svc.ProcessSvc]().Purge(utl.Split(c.Param("ids"), ','))
	r.J(c, true, err)
}

func SaveProcess(c *gin.Context) {
	var process mdl.Process
	err := c.ShouldBind(&process)
	if err != nil {
		r.J(c, err)
		return
	}
	err = s.Use[*svc.ProcessSvc]().Save(process)
	r.J(c, process.Id, err)
}

func ListSteps(c *gin.Context) {
	steps, count, err := s.Use[*svc.ProcessSvc]().ListSteps(cast.ToInt64(c.Param("id")))
	r.J(c, steps, count, err)
}

func RemoveSteps(c *gin.Context) {
	err := s.Use[*svc.ProcessSvc]().RemoveSteps(utl.Split(c.Param("ids"), ','))
	r.J(c, true, err)
}

func SaveStep(c *gin.Context) {
	var step mdl.Step
	err := c.ShouldBind(&step)
	if err != nil {
		r.J(c, err)
		return
	}
	err = s.Use[*svc.ProcessSvc]().SaveStep(step)
	r.J(c, step.Id, err)
}
