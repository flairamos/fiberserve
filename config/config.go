package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	Server      server      `yaml:"server"`
	Database    database    `yaml:"database"`
	Jwt         jwt         `yaml:"jwt"`
	Minio       minio       `yaml:"minio"`
	Redis       redis       `yaml:"redis"`
	Miniprogram miniprogram `yaml:"miniprogram"`
}

type server struct {
	Port          int    `yaml:"port"`
	QrcodeRelease string `yaml:"qrcode_release"`
}

type database struct {
	Driver string `yaml:"driver"`
	User   string `yaml:"user"`
	Passwd string `yaml:"password"`
	Addr   string `yaml:"addr"`
	Port   int    `yaml:"port"`
	DBName string `yaml:"dbname"`
	Zone   string `yaml:"zone"`
}

type jwt struct {
	Key string `yaml:"key"`
}

type minio struct {
	Endpoint     string `yaml:"endpoint"`
	AccessKey    string `yaml:"accesskey"`
	AccessSecret string `yaml:"accesssecret"`
	UseSSL       bool   `yaml:"ssl"`
	Bucket       string `yaml:"bucket"`
}

type redis struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}

type miniprogram struct {
	AppID     string `yaml:"appid"`
	AppSecret string `yaml:"appsecret"`
	OpenApi   string `yaml:"openapi"`
}

var Config config

func Init(filepath string) {
	filePath := filepath + "/config.yaml"
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("读取配置文件失败", err.Error())
		return
	}
	err = yaml.Unmarshal(file, &Config)
	if err != nil {
		log.Println("序列化配置文件失败", err.Error())
		return
	}
}
