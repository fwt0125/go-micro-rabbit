package config

import (
	"fmt"
	"github.com/go-ini/ini"
	"strings"
	"user/model"
)

var (
	Db               string
	DbHost           string
	DbPort           string
	DbUser           string
	DbPassWord       string
	DbName           string
	EtcdAddress      string
	EtcdMicroAddress string
)

func Init() {
	configFile := "./config/config.ini"
	file, err := ini.Load(configFile)
	if err != nil {
		fmt.Println("配置文件路径错误" + configFile)
	}
	LoadEtcdData(file)
	LoadMysqlData(file)
	mysqlPath := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	model.Database(mysqlPath)
}

func LoadEtcdData(file *ini.File) {
	EtcdAddress = file.Section("etcd").Key("Address").String()
	EtcdMicroAddress = file.Section("etcd").Key("MicroAddress").String()
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
}
