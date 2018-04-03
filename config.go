package config

import "github.com/jinzhu/configor"

// Config represent configuration for the app
var Config = struct {
	APPName string `default:"app name"`

	DB struct {
		Name     string
		User     string `default:"root"`
		Password string `default:""`
		Port     uint   `default:"3306"`
	}
}{}

func load() {
	configor.Load(&Config, "tmp/config.yml")
}
