package routers

import (
	"Paymentserver/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MidasController{})
	beego.Router("/payment/createOrder", &controllers.MidasController{}, "get,post:CreateOrder")
	beego.Router("/payment/midasNotify", &controllers.MidasController{}, "get,post:MidasNotify")
}
