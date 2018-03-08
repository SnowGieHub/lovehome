package routers

import (
	"github.com/astaxie/beego"
	"lovehome/controllers"
)

func init() {
	beego.Router("/api/v1.0/houses/:id", &controllers.HouseController{}, "get:GetHouseDetail")
	beego.Router("/api/v1.0/houses", &controllers.HouseSearchController{}, "get:HouseSearch")
	beego.Router("/api/v1.0/user", &controllers.UserSessionController{}, "get:ReadUserSession")
}
