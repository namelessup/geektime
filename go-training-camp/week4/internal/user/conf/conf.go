package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

var Config UserConfig

type UserConfig struct {
	SQLite struct {
		Addr string `yaml:"addr"`
	} `yaml:"sqlite"`
	Grpc struct {
		Addr string `yaml:"addr"`
	} `yaml:"grpc"`
}

func InitUserConfig(name string) error {
	// 从github或者文件读取配置信息
	file, err := ioutil.ReadFile(name)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(file, &Config)
}
