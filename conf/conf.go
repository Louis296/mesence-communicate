package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
	Jwt      Jwt      `yaml:"jwt"`
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

func GetConf() (Conf, error) {
	file, err := ioutil.ReadFile("conf/conf.yaml")
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
