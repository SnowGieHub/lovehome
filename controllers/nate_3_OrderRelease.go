package controllers

import (
	"encoding/json"
	//"fmt"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/cache"
	//_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	"lovehome/models"
	"strconv"
	"time"
)

type OrderController struct {
	beego.Controller
}

//将封装好的返回结构 变成json返回给前段
func (this *OrderController) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *OrderController) OrderRelease() {
	beego.Info("========== /api/v1.0/orders OrderRelease succ ======")

	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)

	var regRequestMap = make(map[string]interface{})
	json.Unmarshal(this.Ctx.Input.RequestBody, &regRequestMap)
	getHid := regRequestMap["house_id"].(string)
	getSd := regRequestMap["start_date"].(string)
	getEd := regRequestMap["end_date"].(string)

	if getHid == "" || getSd == "" || getEd == "" {
		beego.Info("URL中Get后的参数有空值")
		beego.Info("house_id: ", getHid)
		beego.Info("start_date: ", getSd)
		beego.Info("end_date", getEd)
		resp["errno"] = models.RECODE_PARAMERR
		resp["errmsg"] = models.RecodeText(models.RECODE_PARAMERR)
		return
	}

	timeSd, err := time.Parse("2006-01-02", getEd)
	if err != nil {
		beego.Info("sd转为日期时出错")
		resp["errno"] = models.RECODE_SERVERERR
		resp["errmsg"] = models.RecodeText(models.RECODE_SERVERERR)
		return
	}
	timeEd, err := time.Parse("2006-01-02", getEd)
	if err != nil {
		beego.Info("ed转为日期时出错")
		resp["errno"] = models.RECODE_SERVERERR
		resp["errmsg"] = models.RecodeText(models.RECODE_SERVERERR)
		return
	}

	intHid, err := strconv.Atoi(getHid)
	if err != nil {
		beego.Info("aid转为整型时出错")
		resp["errno"] = models.RECODE_SERVERERR
		resp["errmsg"] = models.RecodeText(models.RECODE_SERVERERR)
		return
	}

	if timeSd.After(timeEd) {
		beego.Info("日期有问题")
		resp["errno"] = models.RECODE_PARAMERR
		resp["errmsg"] = models.RecodeText(models.RECODE_PARAMERR)
		return
	}

	o := orm.NewOrm()

	var house models.House

	//house.Id = intHid
	//err := o.Read(&house)
	//if err != nil {
	//	beego.Info("房屋信息数据库读取有问题")
	//	resp["errno"] = models.RECODE_DBERR
	//	resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
	//	return
	//}

	houseTbl := o.QueryTable("user")
	houseTbl.Filter("house__id", intHid).RelatedSel().One(&house)
	if err != nil {
		//返回错误信息给前端
		beego.Info("database err: err ==", err)
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	if house.User.Id == this.GetSession("user_id") {
		beego.Info("不能订购自己的房子 err!")
		resp["errno"] = models.RECODE_REQERR
		resp["errmsg"] = models.RecodeText(models.RECODE_REQERR)
		return
	}

	var newOrder models.OrderHouse
	orderUser := house.User
	timeDiff := timeEd.Sub(timeSd)

	newOrder.User = orderUser
	newOrder.House = &house
	newOrder.Begin_date = timeSd
	newOrder.End_date = timeEd
	newOrder.Days = int(timeDiff)
	newOrder.House_price = house.Price
	newOrder.Amount = house.Deposit
	newOrder.Ctime = time.Now()
	id, err := o.Insert(&newOrder)
	if err == nil {
		beego.Info("房屋信息数据库插入有问题")
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	resp["order_id"] = strconv.Itoa(int(id))
	return
}
