package main

import (
	"Centreserver/client"
	"Centreserver/gm"
	"galaxy"
	"galaxy/utils"
)

func main() {
	defer utils.Stack()

	galaxy.GxService().Run()
	client.ClientManager()
	gm.GMWebInit(":8088")
	galaxy.GxService().Wait()
}
