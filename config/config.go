package config

import (
	"github.com/BurntSushi/toml"
)

type ossInfo struct {
	EndPoint        string
	AccessKeyID     string
	AccessKeySecret string
	Bucket          string
}

type Config struct {
	Oss ossInfo
}

func InitConfig() (cgg Config, err error) {
	var cg Config
	_, err = toml.DecodeFile("./config.toml", &cg)
	return cg, err
}
