package main

import (
	"Gateserver/client"
	"galaxy"
	"galaxy/utils"
)

func main() {
	defer utils.Stack()

	galaxy.GxService().Run()
	client.ClientManager()
	galaxy.GxService().Wait()

}
