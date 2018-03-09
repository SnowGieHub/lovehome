package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"lovehome/models"
)

type HouseController struct {
	beego.Controller
}

//func (this *HousesIndexController) RetData(resp interface{}) {
//	this.Data["json"] = resp
//	this.ServeJSON()
//}

func (this *HouseController) GetHouseDetail() {
	beego.Info("GetHouseDetail .....")
	house_id := this.Ctx.Input.Param(":id")
	var myHouse []models.House
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	o := orm.NewOrm()
	qs := o.QueryTable("house")
	err := qs.Filter("id", house_id).One(&myHouse)
	if err != nil {
		if err == orm.ErrNoRows {
			resp["errno"] = models.RECODE_NODATA
			resp["errmsg"] = models.RecodeText(models.RECODE_NODATA)
			return
		} else {
			resp["errno"] = models.RECODE_DBERR
			resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
			return
		}
	} else {
		resp["errno"] = 0
		resp["errmsg"] = "成功"
		myData := make(map[string]interface{})

		myData["house"] = myHouse
		myData["user_id"] = this.GetSession("user_id")
		resp["data"] = myData
		return
	}
}
