package config

import (
	"fmt"
	"github.com/go-ini/ini"
)

var (
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
}

func LoadEtcdData(file *ini.File) {
	EtcdAddress = file.Section("etcd").Key("Address").String()
	EtcdMicroAddress = file.Section("etcd").Key("MicroAddress").String()
}
