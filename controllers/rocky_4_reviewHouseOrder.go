package controllers

import (
	_ "encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"lovehome/models"
	_ "path"
)

type RevHouseOrdersController struct {
	beego.Controller
}

//将从数据库中读取到的字段信息转换成 前端需要的json 格式
func (this *RevHouseOrdersController) ArraryOrder(orderhouse []*models.OrderHouse) (myorders []interface{}) {
	for _, value := range orderhouse {
		order := make(map[string]interface{})

		order["amount"] = value.Amount
		order["comment"] = value.Comment
		order["ctime"] = string(value.Ctime.Format("2006-01-02 15:04:05"))
		order["days"] = value.Days
		order["end_date"] = string(value.End_date.Format("2006-01-02 15:04:05"))
		order["img_url"] = value.House.Index_image_url
		order["order_id"] = value.Id
		order["start_date"] = string(value.Begin_date.Format("2006-01-02 15:04:05"))
		order["status"] = value.Status
		order["title"] = value.House.Title
		myorders = append(myorders, order)
		beego.Info("............", order)
	}

	return myorders

}

//将封装好的返回结构 变成json返回给前端
func (this *RevHouseOrdersController) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *RevHouseOrdersController) ReviewHouseOrders() {
	//获取session  中name的信息
	name_info := this.GetSession("name")
	id_info := this.GetSession("user_id")

	beego.Info("session id name = ", name_info)

	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	//获取url 中 是顾客还是房东还是空
	kind_customer := this.GetString("role")
	if kind_customer == "" {
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}

	defer this.RetData(resp)

	//创建查询数据库的句柄
	var orderhouse []*models.OrderHouse
	var house []*models.House

	o := orm.NewOrm()
	beego.Info("kind_customer = ", kind_customer)
	orders := make(map[string]interface{}) //
	myorders := []interface{}{}            //用于拼接json 数组

	//判断是房东还是租客
	if kind_customer == "custom" {
		//查询数据库的订单表
		qs := o.QueryTable("order_house")
		beego.Info(name_info, "-------------")
		num, err := qs.Filter("user__name", name_info).All(&orderhouse)
		beego.Info("num ----------", num)
		if err != nil {

			resp["errno"] = models.RECODE_DBERR
			resp["errmsp"] = models.RecodeText(models.RECODE_DBERR)
			return
		}

		myorders = this.ArraryOrder(orderhouse) //遍历出orderhouse 中的条数，并追加

	} else if kind_customer == "landlord" {

		qs := o.QueryTable("house")
		beego.Info("id name ", id_info)
		num, err := qs.Filter("user__id", id_info).All(&house) //查出user——id下的所有房源
		beego.Info("name_info ", name_info)
		if err != nil {
			resp["errno"] = models.RECODE_DBERR
			resp["errmsp"] = models.RecodeText(models.RECODE_DBERR)
			return

		}
		beego.Info("num ------ = ", num)

		for _, value := range house {
			//拿到house的id信息

			beego.Info("value :=  ", value.Id)
			qs := o.QueryTable("order_house")
			num, err := qs.Filter("user_id", id_info).All(&orderhouse)
			beego.Info("num order house ----------", num)
			if err != nil {

				resp["errno"] = models.RECODE_DBERR
				resp["errmsp"] = models.RecodeText(models.RECODE_DBERR)
				return
			}

			//然后根据房屋的id 查询orderhouse 表 中与用户id和house  id 相同的条数
			myorders = this.ArraryOrder(orderhouse)

		}

	}
	fmt.Printf("%s+v\n", myorders)

	orders["orders"] = myorders

	//将数据返回给前端
	resp["data"] = orders

}
