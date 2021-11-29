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

func InitConfig() (cg Config, err error) {
	_, err = toml.DecodeFile("./config.toml", &cg)
	return cg, err
}
