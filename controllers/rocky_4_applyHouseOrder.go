package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"lovehome/models"
	"strconv"
)

type ApplyOrderController struct {
	beego.Controller
}

func (this *ApplyOrderController) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *ApplyOrderController) ErrReturn() {

}

func (this *ApplyOrderController) AplyOrders() {

	beego.Info("接单已调用------------------------")

	resp := make(map[string]interface{})
	//初始化resp data
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer this.RetData(resp)

	aaction := make(map[string]interface{})

	json.Unmarshal(this.Ctx.Input.RequestBody, &aaction)
	//定义数据库查询句柄
	o := orm.NewOrm()
	var orderhouse models.OrderHouse
	//得到用户id
	user_info := this.GetSession("user_id")

	//获取url中的订单号
	orderNum := this.Ctx.Input.Param(":id")

	//查询订单表 确认订单状态是否是WAIT_ACCEPT
	beego.Info("用户id 为 ", user_info)
	beego.Info("url 中的订单编号为", orderNum)

	//获取put的数据动作
	action := aaction["action"]

	beego.Info("action is ----------------", aaction["action"])
	if action == "" {
		resp["errno"] = models.RECODE_NODATA
		resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
		return
	}
	qs := o.QueryTable("order_house")
	//2.创建查询订单的存储
	qs.Filter("id", orderNum).One(&orderhouse)

	beego.Info("数据库中的数据为----------", orderhouse)

	if action == "accept" { //accept 动作

		if orderhouse.Status == "WAIT_ACCEPT" {
			beego.Info("是等待接单状态-------------")

			//核对当前用户id 和从表中查出来的id 是否一致
			if orderhouse.User.Id == user_info {

				beego.Info("id 一致可以继续操作--------")
				//更改订单状态为ACCEPT
				orderid, _ := strconv.Atoi(orderNum)

				beego.Info("000000000000000", orderid)
				order := models.OrderHouse{Id: orderid}

				order.Status = "ACCEPT"

				beego.Info("order ====", order)
				//插入到数据库

				o_num, err := o.Update(&order, "status")
				if err != nil {
					resp["errno"] = models.RECODE_DBERR
					resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
					return

				}
				beego.Info("order num ====== ", o_num)

				//返回正确json

				resp["errno"] = models.RECODE_OK
				resp["errmsg"] = models.RecodeText(models.RECODE_OK)
			}
		}
	} else if action == "reject" { //reject 动作处理
		//如果不一致就更改数据库状态为REJECT 并将从url中获取reson

		orderid, _ := strconv.Atoi(orderNum)
		order := models.OrderHouse{Id: orderid}
		order.Status = "REJECT"
		//插入到数据库
		o_num, err := o.Update(&order, "status")
		if err != nil {
			resp["errno"] = models.RECODE_DBERR
			resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
			return

		}
		beego.Info("order num ====== ", o_num)
		//从数据库中读取评价信息

		//从url 中获取 reson 字段
		rreason := make(map[string]interface{})

		json.Unmarshal(this.Ctx.Input.RequestBody, &rreason)
		//追加reason 信息至comment 中

		beego.Info("原有数据库表中的评价信息", orderhouse.Comment)
		reason_info := orderhouse.Comment + " " + rreason["reason"].(string)

		commentOrder := models.OrderHouse{Id: orderid}
		commentOrder.Comment = reason_info
		//将comment信息更新到数据库中
		beego.Info("评价信息追加为  ", commentOrder)
		o_num, err = o.Update(&commentOrder, "comment")
		if err != nil {
			resp["errno"] = models.RECODE_DBERR
			resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
			return

		}

		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		return

	}
}
