package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	Loadbalancer struct {
		Servers []struct {
			Id        int    `yaml:"id"`
			ServerURI string `yaml:"serverURI"`
		}
	}
}

func (c *Config) parseConfig() Config {
	filename, _ := filepath.Abs("server-list.yaml")

	yamlFile, err := ioutil.ReadFile(filename)
	checkErr(err)

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	checkErr(err)

	return config

}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func GetServerList() []string {
	var serverList []string
	var c Config
	parsedYaml := c.parseConfig()
	serversArr := parsedYaml.Loadbalancer.Servers

	for i := 0; i < len(serversArr); i++ {
		serverInfo := serversArr[i]
		serverList = append(serverList, serverInfo.ServerURI)
	}

	return serverList

}
