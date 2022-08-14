package config

import (
	"fmt"
	"github.com/go-ini/ini"
	"strings"
	"task/model"
)

var (
	Db               	string
	DbHost           	string
	DbPort           	string
	DbUser           	string
	DbPassWord       	string
	DbName           	string
	EtcdAddress      	string
	EtcdMicroAddress 	string
	RabbitMQ			string
	RabbitMQUser		string
	RabbitMQPassword	string
	RabbitMQHost		string
	RabbitMQPort		string
)

func Init() {
	configFile := "./config/config.ini"
	file, err := ini.Load(configFile)
	if err != nil {
		fmt.Println("配置文件路径错误" + configFile)
	}
	LoadMysqlData(file)
	mysqlPath := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	model.Database(mysqlPath)

	LoadRabbitMQData(file)
	rabbitMQPath := strings.Join([]string{RabbitMQ, "://", RabbitMQUser, ":", RabbitMQPassword, "@",RabbitMQHost, ":",RabbitMQPort, "/"}, "")
	model.RabbitMQ(rabbitMQPath)

}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbName = file.Section("mysql").Key("DbName").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
}


func LoadRabbitMQData(file *ini.File) {
	RabbitMQ = file.Section("rabbitmq").Key("RabbitMQ").String()
	RabbitMQUser = file.Section("rabbitmq").Key("RabbitMQUser").String()
	RabbitMQPassword = file.Section("rabbitmq").Key("RabbitMQPassword").String()
	RabbitMQHost = file.Section("rabbitmq").Key("RabbitMQHost").String()
	RabbitMQPort = file.Section("rabbitmq").Key("RabbitMQPort").String()
}