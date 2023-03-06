package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

var (
	configFileName string = "config.yml"
)

type config struct {
	Loadbalancer struct {
		Servers []struct {
			Id        int    `yaml:"id"`
			ServerURI string `yaml:"serverURI"`
		}
	}
}

func (c *config) parseConfig() config {
	filename, _ := filepath.Abs(configFileName)

	yamlFile, err := ioutil.ReadFile(filename)
	checkErr(err)

	var config config

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
	var c config
	parsedYaml := c.parseConfig()
	serversArr := parsedYaml.Loadbalancer.Servers

	for i := 0; i < len(serversArr); i++ {
		serverInfo := serversArr[i]
		serverList = append(serverList, serverInfo.ServerURI)
	}

	return serverList

}
