package main

import (
	"bryanyi.com/config"
	"bryanyi.com/loadbalancer"
	"fmt"
)

func main() {
	serverList := config.GetServerList()
	serverCount := len(serverList)

	fmt.Println("server count: ", serverCount, "servers array: ", serverList)

	if serverCount > 5 {
		loadbalancer.MakeLoadBalancer(serverCount, serverList)
	}

}
