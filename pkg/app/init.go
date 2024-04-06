package app

import (
	"pkg/mysql"
	"pkg/nacos"
)

const NACOSADDRESS = "127.0.0.1"
const NACOSPORT = 8848
const SERVICEGROUP = "DEFAULT_GROUP"
const SERVICENAME = "zg6"

func Init(apps ...string) error {
	err := nacos.GetClient(NACOSADDRESS, NACOSPORT)

	for _, app := range apps {
		switch app {
		case "mysql":
			err = mysql.InitMysql(SERVICEGROUP, SERVICENAME)
			if err != nil {
				return err
			}

		}
	}

	return err
}
