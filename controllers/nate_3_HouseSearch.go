package controllers

import (
	//"encoding/json"
	//"fmt"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/cache"
	//_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	"lovehome/models"
	"strconv"
	"time"
)

type HouseSearchController struct {
	beego.Controller
}

//将封装好的返回结构 变成json返回给前段
func (this *HouseSearchController) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *HouseSearchController) HouseSearch() {
	beego.Info("========== /api/v1.0/houses/ HouseSearch  succ ======")

	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)

	getAid := this.GetString("aid")
	getSd := this.GetString("sd")
	getEd := this.GetString("ed")
	getSk := this.GetString("sk")
	getP := this.GetString("p")
	if getAid == "" || getSd == "" ||
		getEd == "" || getSk == "" || getP == "" {
		beego.Info("URL中Get后的参数有空值")
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
	intP, err := strconv.Atoi(getP)
	if err != nil {
		beego.Info("p转为整型时出错")
		resp["errno"] = models.RECODE_SERVERERR
		resp["errmsg"] = models.RecodeText(models.RECODE_SERVERERR)
		return
	}

	intAid, err := strconv.Atoi(getAid)
	if err != nil {
		beego.Info("aid转为整型时出错")
		resp["errno"] = models.RECODE_SERVERERR
		resp["errmsg"] = models.RecodeText(models.RECODE_SERVERERR)
		return
	}

	if timeSd.After(timeEd) || intP <= 0 {
		beego.Info("日期有问题或p值小于等于0")
		resp["errno"] = models.RECODE_PARAMERR
		resp["errmsg"] = models.RecodeText(models.RECODE_PARAMERR)
		return
	}

	//如果Redis中没有，读取数据库
	o := orm.NewOrm()

	var houses []models.House
	dataMap := make(map[string]interface{})

	qs := o.QueryTable("house")
	_, err = qs.Filter("area__id", intAid).All(&houses)
	if err != nil {
		//返回错误信息给前端
		beego.Info("database err: err ==", err)
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	dataMap["houses"] = houses
	resp["data"] = dataMap
	//把数据发挥前端
}
