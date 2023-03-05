package main

import (
	"fmt"

	// "bryanyi.com/loadbalancer"
	"github.com/spf13/viper"
)

func main() {
	// loadbalancer.MakeLoadBalancer(3)

	vi := viper.New()
	vi.SetConfigFile("application.yml")
	vi.ReadInConfig()
	server1Port := vi.GetString("loadbalancer.servers.server1.uri")
	server2Port := vi.GetString("loadbalancer.servers.server2.uri")
	server3Port := vi.GetString("loadbalancer.servers.server3.uri")
	server4Port := vi.GetString("loadbalancer.servers.server4.uri")

	serverArray := []string{}
	serverArray = append(serverArray, server1Port)
	serverArray = append(serverArray, server2Port)
	serverArray = append(serverArray, server3Port)
	serverArray = append(serverArray, server4Port)

	fmt.Println(serverArray)

}
