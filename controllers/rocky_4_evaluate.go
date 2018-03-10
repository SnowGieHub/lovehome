package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"lovehome/models"
	_ "path"
	"strconv"
)

type EvaluateController struct {
	beego.Controller
}

//将封装好的返回结构 变成json返回给前端
func (this *EvaluateController) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *EvaluateController) EvaluateMng() {

	beego.Info("evaluate mng 被调用=-------------")

	//获取session  中name的信息
	resp := make(map[string]interface{})

	//初始化resp
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	order_info := make(map[string]interface{})
	json.Unmarshal(this.Ctx.Input.RequestBody, &order_info)

	defer this.RetData(resp)

	//获取put 请求中的订单编号
	order_id := order_info["order_id"]
	comment := order_info["comment"]

	//创建数据库句柄, 注意评价是追加的，所以要先查一下
	o := orm.NewOrm()
	//创建接收order_house的变量
	var oneOrder models.OrderHouse
	qs := o.QueryTable("order_house")
	err := qs.Filter("id", order_id).One(&oneOrder) //将查到的信息取出
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	comment_info := oneOrder.Comment + " " + comment.(string)
	id_user, _ := strconv.Atoi(order_id.(string)) //将interface接口的数据转换成int 类型
	order := models.OrderHouse{Id: id_user}
	order.Comment = comment_info
	o_num, err := o.Update(&order, "comment")
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return

	}
	beego.Info("order num ====== ", o_num)
	return
}
