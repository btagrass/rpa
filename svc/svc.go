package svc

import (
	s "github.com/btagrass/gobiz/svc"
	"github.com/btagrass/gobiz/sys/svc"
	"github.com/sirupsen/logrus"
)

func init() {
	svc.Init()
	err := s.Migrate(
		"INSERT IGNORE INTO sys_resource VALUES (300000000000002, '2023-01-29 00:00:00.000', NULL, NULL, 0, 'BusinessSystem', 1, 'Operation', '/mgt', NULL, 1)",
		"INSERT IGNORE INTO sys_resource VALUES (300000000000201, '2023-01-29 00:00:00.000', NULL, NULL, 300000000000002, 'ProcessManagement', 1, 'Monitor', '/mgt/processes', NULL, 1)",
		"INSERT IGNORE INTO sys_resource VALUES (300000000020101, '2023-01-29 00:00:00.000', NULL, NULL, 300000000000201, 'Query', 2, '', '/mgt/processes', 'GET', 1)",
		"INSERT IGNORE INTO sys_resource VALUES (300000000020102, '2023-01-29 00:00:00.000', NULL, NULL, 300000000000201, 'Delete', 2, '', '/mgt/processes/:ids', 'DELETE', 2)",
		"INSERT IGNORE INTO sys_resource VALUES (300000000020103, '2023-01-29 00:00:00.000', NULL, NULL, 300000000000201, 'Edit', 2, '', '/mgt/processes/:id', 'GET', 3)",
		"INSERT IGNORE INTO sys_resource VALUES (300000000020104, '2023-01-29 00:00:00.000', NULL, NULL, 300000000000201, 'Save', 2, '', '/mgt/processes', 'POST', 4)",
		"INSERT IGNORE INTO sys_dict VALUES (300000002010101, '2023-01-29 00:00:00.000', NULL, NULL, 'ProcessType', 1, 'Process', 1)",
		"INSERT IGNORE INTO sys_dict VALUES (300000002010111, '2023-01-29 00:00:00.000', NULL, NULL, 'ProcessType', 11, 'Wait', 11)",
		"INSERT IGNORE INTO sys_dict VALUES (300000002010121, '2023-01-29 00:00:00.000', NULL, NULL, 'ProcessType', 21, '[Web]Goto', 21)",
		"INSERT IGNORE INTO sys_dict VALUES (300000002010122, '2023-01-29 00:00:00.000', NULL, NULL, 'ProcessType', 22, '[Web]Attr', 22)",
		"INSERT IGNORE INTO sys_dict VALUES (300000002010123, '2023-01-29 00:00:00.000', NULL, NULL, 'ProcessType', 23, '[Web]Click', 23)",
		"INSERT IGNORE INTO sys_dict VALUES (300000002010124, '2023-01-29 00:00:00.000', NULL, NULL, 'ProcessType', 24, '[Web]Input', 24)",
		"INSERT IGNORE INTO sys_dict VALUES (300000002010125, '2023-01-29 00:00:00.000', NULL, NULL, 'ProcessType', 25, '[Web]Text', 25)",
	)
	if err != nil {
		logrus.Fatal(err)
	}
	s.Inject(NewProcessSvc)
}
