package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"rpa/api"
	_ "rpa/job"
	"rpa/mgt"
	_ "net/http/pprof"

	"github.com/btagrass/gobiz/cmd"
	"github.com/btagrass/gobiz/htp"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

//go:embed web/dist
var dist embed.FS

func main() {
	cmd.Execute(
		"rpa",
		"TreeRpa",
		cmd.Install,
		cmd.Start,
		cmd.Status,
		cmd.Stop,
		cmd.Uninstall,
		&cobra.Command{
			Use:   "run",
			Short: "Run",
			Run: func(c *cobra.Command, args []string) {
				// process := mdl.Process{
				// 	Children: []mdl.Process{
				// 		{
				// 			Type:    "navigate",
				// 			Content: "{\"uri\":\"http://www.baidu.com\",\"visible\":true,\"timeout\":\"3s\"}",
				// 		},
				// 		{
				// 			Type:    "input",
				// 			Content: "{\"selector\":\"#kw\",\"value\":\"abc\"}",
				// 		},
				// 		{
				// 			Type:    "click",
				// 			Content: "{\"selector\":\"#su\"}",
				// 		},
				// 		{
				// 			Type:    "wait",
				// 			Content: "5s",
				// 		},
				// 		{
				// 			Type:    "readExcel",
				// 			Content: "{\"filename\":\"/home/btag/文档/course.xlsx\",\"sheet\":\"sheet1\",\"cols\":\"1,2\"}",
				// 			VarName: "表格",
				// 		},
				// 	},
				// }
				// vars, err := s.Use[*svc.ProcessSvc]().Run(process)
				// fmt.Println(vars, err)

				group := &errgroup.Group{}
				// Auth
				group.Go(func() error {
					expirationDate, _ := time.Parse(time.DateOnly, "2025-04-13")
					if time.Now().After(expirationDate) {
						panic("系统内部异常")
					}
					return nil
				})
				// Api
				api := &http.Server{
					Addr:    fmt.Sprintf(":%d", htp.Port),
					Handler: api.Api(),
				}
				group.Go(api.ListenAndServe)
				// Mgt
				mgt := &http.Server{
					Addr:    fmt.Sprintf(":%d", htp.Port+1),
					Handler: mgt.Mgt(),
				}
				group.Go(mgt.ListenAndServe)
				// Web
				dist, _ := fs.Sub(dist, "web/dist")
				engine := gin.Default()
				engine.StaticFS("/", http.FS(dist))
				web := &http.Server{
					Addr:    fmt.Sprintf(":%d", htp.Port+2),
					Handler: engine,
				}
				group.Go(web.ListenAndServe)
				err := group.Wait()
				if err != nil {
					logrus.Fatal(err)
				}
			},
		},
	)
}
