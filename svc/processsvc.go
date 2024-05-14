package svc

import (
	"fmt"
	"rpa/mdl"
	"time"

	"github.com/btagrass/gobiz/svc"
	"github.com/btagrass/gobiz/uat"
	"github.com/btagrass/gobiz/utl"
	"github.com/samber/do"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"github.com/xuri/excelize/v2"
)

type ProcessSvc struct {
	*svc.DataSvc[mdl.Process]
	stepDataSvc *svc.DataSvc[mdl.Step]
}

func NewProcessSvc(i *do.Injector) (*ProcessSvc, error) {
	return &ProcessSvc{
		DataSvc:     svc.NewDataSvc[mdl.Process]("rpa:processes", new(mdl.Step)),
		stepDataSvc: svc.NewDataSvc[mdl.Step]("rpa:steps", new(mdl.Step)),
	}, nil
}

func (s *ProcessSvc) Run(id int64) (map[string]any, error) {
	variables := make(map[string]any)
	var browser *uat.Browser
	process, err := s.Get(id)
	if err != nil {
		return variables, err
	}
	for _, ps := range process.Steps {
		switch ps.Type {
		// Common
		case "Wait":
			time.Sleep(cast.ToDuration(ps.Duration))
		// Excel
		case "[Excel]Read":
			cells := make([]map[int]any, 0)
			sheet := gjson.Get(ps.Name, "sheet").String()
			if sheet == "" {
				sheet = "Sheet1"
			}
			cols := utl.Split(gjson.Get(ps.Name, "cols").String(), ',')
			if ps.Variable == "" {
				ps.Variable = fmt.Sprintf("%s_%s", ps.Name, sheet)
			}
			file, err := excelize.OpenFile(ps.Name)
			if err != nil {
				if ps.Error == "continue" {
					continue
				} else {
					return variables, err
				}
			}
			defer file.Close()
			rows, err := file.GetRows(sheet)
			if err != nil {
				if ps.Error == "continue" {
					continue
				} else {
					return variables, err
				}
			}
			for _, r := range rows {
				colums := make(map[int]any)
				for j, c := range r {
					if len(cols) == 0 || lo.Contains[string](cols, cast.ToString(j+1)) {
						colums[j+1] = c
					}
				}
				cells = append(cells, colums)
			}
			variables[ps.Variable] = cells
		case "[Excel]Write":
			sheet := gjson.Get(ps.Name, "sheet").String()
			cols := utl.Split(gjson.Get(ps.Name, "cols").String(), ',')
			file := excelize.NewFile()
			defer file.Close()
			rows, ok := variables[ps.Variable].([]map[string]any)
			if ok {
				for j, c := range cols {
					cell, err := excelize.CoordinatesToCellName(j+1, 1)
					if err != nil {
						if ps.Error == "continue" {
							continue
						} else {
							return variables, err
						}
					}
					err = file.SetCellValue(sheet, cell, c)
					if err != nil {
						if ps.Error == "continue" {
							continue
						} else {
							return variables, err
						}
					}
				}
				for i, r := range rows {
					for j, c := range cols {
						cell, err := excelize.CoordinatesToCellName(j+1, i+2)
						if err != nil {
							if ps.Error == "continue" {
								continue
							} else {
								return variables, err
							}
						}
						err = file.SetCellValue(sheet, cell, r[c])
						if err != nil {
							if ps.Error == "continue" {
								continue
							} else {
								return variables, err
							}
						}
					}
				}
				err = file.SaveAs(ps.Name)
				if err != nil {
					if ps.Error == "continue" {
						continue
					} else {
						return variables, err
					}
				}
			}
		// Web
		case "[Web]Attr":
			if ps.Variable == "" {
				ps.Variable = fmt.Sprintf("%s_%s", ps.Name, ps.Value)
			}
			value, err := browser.Attr(ps.Name, ps.Visible, ps.Value)
			if err != nil {
				if ps.Error == "continue" {
					continue
				} else {
					return variables, err
				}
			}
			variables[ps.Variable] = value
		case "[Web]Click":
			err = browser.Click(ps.Name, ps.Visible)
			if err != nil {
				if ps.Error == "continue" {
					continue
				} else {
					return variables, err
				}
			}
		case "[Web]Goto":
			if browser == nil {
				browser, err = uat.NewBrowser(ps.Name, cast.ToBool(ps.Value), "", map[string]string{})
				if err != nil {
					return variables, err
				}
			}
		case "[Web]Input":
			err = browser.Input(ps.Name, ps.Visible, ps.Value)
			if err != nil {
				if ps.Error == "continue" {
					continue
				} else {
					return variables, err
				}
			}
		case "[Web]Text":
			if ps.Variable == "" {
				ps.Variable = fmt.Sprintf("%s_text", ps.Name)
			}
			text, err := browser.Text(ps.Name, ps.Visible)
			if err != nil {
				if ps.Error == "continue" {
					continue
				} else {
					return variables, err
				}
			}
			variables[ps.Variable] = text
		}
	}
	return variables, nil
}

func (s *ProcessSvc) ListSteps(processId int64) ([]mdl.Step, int64, error) {
	return s.stepDataSvc.List("process_id = ?", processId)
}

func (s *ProcessSvc) RemoveSteps(ids []string) error {
	return s.stepDataSvc.Purge(ids)
}

func (s *ProcessSvc) SaveStep(step mdl.Step) error {
	return s.stepDataSvc.Save(step)
}
