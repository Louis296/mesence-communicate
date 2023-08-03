package conf

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Conf struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Jwt      Jwt      `yaml:"jwt"`
	MongoDB  MongoDB  `yaml:"mongodb"`
}
type Server struct {
	Port int `yaml:"port"`
}
type Database struct {
	URL          string `yaml:"url"`
	UserName     string `yaml:"userName"`
	Password     string `yaml:"password"`
	DatabaseName string `yaml:"databaseName"`
	MaxConn      int    `yaml:"maxConn"`
	MaxOpen      int    `yaml:"maxOpen"`
}
type Jwt struct {
	Secret string `yaml:"secret"`
}
type MongoDB struct {
	Url string `yaml:"url"`
}

func GetConf() (Conf, error) {
	file, err := os.ReadFile("conf/conf.yaml")
	if err != nil {
		return Conf{}, err
	}
	var ans Conf
	err = yaml.UnmarshalStrict(file, &ans)
	if err != nil {
		return Conf{}, err
	}
	return ans, nil
}
