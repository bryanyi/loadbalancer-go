package main

import (
	"bryanyi.com/config"
	"bryanyi.com/loadbalancer"
)

func main() {
	serverList := config.GetServerList()
	serverCount := len(serverList)

	if serverCount > 1 {
  loadbalancer.MakeLoadBalancer(serverCount, serverList)
	}

}
