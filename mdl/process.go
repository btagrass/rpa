package mdl

import "github.com/btagrass/gobiz/mdl"

type Process struct {
	mdl.Mdl
	Name  string `gorm:"" json:"name"`
	Steps []Step `gorm:"foreignKey:ProcessId" json:"steps"`
}

type Step struct {
	mdl.Mdl
	ProcessId int64  `gorm:"not null" json:"processId"`
	ParentId  int64  `gorm:"not null" json:"parentId"`
	Type      string `gorm:"" json:"type"`
	Name      string `gorm:"" json:"name"`
	Visible   bool   `gorm:"" json:"visible"`
	Value     string `gorm:"" json:"value"`
	Duration  int    `gorm:"" json:"duration"`
	Variable  string `gorm:"" json:"variable"`
	Error     string `gorm:"" json:"error"`
	Children  []Step `gorm:"foreignKey:ParentId" json:"children"`
}
