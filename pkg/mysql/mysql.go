package mysql

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"pkg/nacos"
)

var Db *gorm.DB

// 连接mysql
type MysqlConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}
type Val struct {
	Mysql MysqlConfig `yaml:"mysql"`
}

var MysqlConfigVal Val

func GetMysqlConfig(serviceGroup, serviceName string) error {
	content, err := nacos.GetConfig(serviceGroup, serviceName)
	fmt.Println(7777777777777)
	fmt.Println(content)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(content), &MysqlConfigVal)
	if err != nil {
		return err
	}

	return nil
}

func InitMysql(serviceGroup, serviceName string) error {
	err := GetMysqlConfig(serviceGroup, serviceName)
	if err != nil {
		log.Println("failed to get mysql config")
		return err
	}
	mConfig := MysqlConfigVal.Mysql
	fmt.Println(mConfig)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mConfig.Username, mConfig.Password, mConfig.Host, mConfig.Port, mConfig.Database)
	//dsn := "root:root@tcp(127.0.0.1:3306)/zg6?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}
	return nil
}
