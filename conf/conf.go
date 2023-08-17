package conf

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Conf struct {
	Server  Server  `yaml:"server"`
	MySQL   MySQL   `yaml:"mysql"`
	Jwt     Jwt     `yaml:"jwt"`
	MongoDB MongoDB `yaml:"mongodb"`
	Kafka   Kafka   `yaml:"kafka"`
	Redis   Redis   `yaml:"redis"`
}
type Server struct {
	Port int `yaml:"port"`
}
type MySQL struct {
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
	Url      string `yaml:"url"`
	Database string `yaml:"database"`
}
type Kafka struct {
	Url   string `yaml:"url"`
	Topic string `yaml:"topic"`
}
type Redis struct {
	Url      string `yaml:"url"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
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
