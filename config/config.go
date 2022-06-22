package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type App struct {
	Address      string
	Static       string
	Log          string
	Time         string
	OssMaxKeys   int
	OnLine       int
	IsShowOnLine int
	OssUrl       string
	Limit        int
}
type Redis struct {
	Address string
	Auth    string
	Key     string
}

type Database struct {
	Driver   string
	Address  string
	Database string
	User     string
	Password string
}

type Configuration struct {
	App   App
	Redis Redis
	Db    Database
}

var config *Configuration
var once sync.Once

// 通过单例模式初始化全局配置
func LoadConfig() *Configuration {
	once.Do(func() {
		file, err := os.Open("config.json")
		if err != nil {
			log.Fatalln("Cannot open config file", err)
		}
		decoder := json.NewDecoder(file)
		config = &Configuration{}
		err = decoder.Decode(config)
		if err != nil {
			log.Fatalln("Cannot get configuration from file", err)
		}
	})
	return config
}
