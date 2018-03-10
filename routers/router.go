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
	//李雪楠
	beego.Router("/api/v1.0/user/name", &controllers.UserController{}, "Put:PutName")
	beego.Router("api/v1.0/user/auth", &controllers.UserController{}, "get:AuthGet")
	beego.Router("api/v1.0/user/auth", &controllers.UserController{}, "post:AuthPost")

	// 1
	// 李楠-user-0
	beego.Router("/api/v1.0/user", &controllers.UserSessionController{}, "get:ReadUserSession")
	beego.Router("/api/v1.0/houses/:id", &controllers.HouseController{}, "get:GetHouseDetail")
	beego.Router("/api/v1.0/houses", &controllers.HouseSearchController{}, "get:HouseSearch")
	beego.Router("/api/v1.0/orders", &controllers.OrderController{}, "post:OrderRelease")

	//2
	//龚文斌
	beego.Router("/api/v1.0/user/houses", &controllers.HousesController{}, "get:GetHouseInfo")
	beego.Router("/api/v1.0/houses", &controllers.HousesController{}, "post:ReleaseHouseInfo")
	beego.Router("/api/v1.0/houses/:id/images", &controllers.UpImageControllers{}, "post:UpHouseImage")

	//3
	//唐有利
	//modify by rocky 房屋订单查询 //请求查看房东/租客订单信息  订单插入
	beego.Router("/api/v1.0/user/orders", &controllers.RevHouseOrdersController{}, "get:ReviewHouseOrders")

	//modify by rocky 房东同意/拒绝订单
	beego.Router("/api/v1.0/orders/:id/status", &controllers.ApplyOrderController{}, "put:AplyOrders")

	//modify by rocky  评价信息
	beego.Router("/api/v1.0/orders/:id/comment", &controllers.EvaluateController{}, "put:EvaluateMng")
}
