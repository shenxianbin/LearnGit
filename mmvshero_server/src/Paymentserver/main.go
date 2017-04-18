package main

import (
	"Paymentserver/models"
	_ "Paymentserver/routers"
	"github.com/astaxie/beego"
)

func main() {
	models.LoadConfig()
	beego.Run()
}
