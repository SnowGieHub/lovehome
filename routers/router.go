package routers

import (
	"github.com/astaxie/beego"
	"lovehome/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//请求地域信息
	beego.Router("/api/v1.0/areas", &controllers.AreaController{}, "get:GetAreaInfo")
	//session
	beego.Router("/api/v1.0/session", &controllers.SessionController{}, "get:GetSessionName;delete:DelSessionName")
	//登陆
	beego.Router("api/v1.0/sessions", &controllers.UserController{}, "post:Login")

	//house/index  房屋首页列表
	beego.Router("/api/v1.0/houses/index", &controllers.HousesIndexController{}, "get:GetHousesIndex")

	beego.Router("/api/v1.0/user/avatar", &controllers.UserController{}, "post:UploadAvatar")
	beego.Router("/api/v1.0/users", &controllers.UserController{}, "post:Reg")

	// 0
	//李雪楠-0-请求更新用户名
	beego.Router("/api/v1.0/user/name", &controllers.UserController{}, "Put:PutName")

	//李雪楠-1-实名制请求-get
	beego.Router("api/v1.0/user/auth", &controllers.UserController{}, "get:AuthGet")

	//李雪楠-2-实名制查询-post
	//	beego.Router("api/v1.0/user/auth", &controllers.UserController{}, "post:AuthPost")

	// 1

	// 李楠-user-0
	beego.Router("/api/v1.0/user", &controllers.UserSessionController{}, "get:ReadUserSession")
	//beego.Router("/api/v1.0/houses/:id", &controllers.HouseController{}, "get:GetHouseDetail")
	//beego.Router("/api/v1.0/houses", &controllers.HouseSearchController{}, "get:HouseSearch")

}
