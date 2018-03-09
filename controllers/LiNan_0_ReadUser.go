package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"lovehome/models"
)

type UserSessionController struct {
	beego.Controller
}

func (this *UserSessionController) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *UserSessionController) ReadUserSession() {
	beego.Info("========== /api/v1.0/user/ ReadUserSession  succ ======")

	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)

	user_id := this.GetSession("user_id")
	if user_id == nil {
		resp["errno"] = models.RECODE_SESSIONERR
		resp["errmsg"] = models.RecodeText(models.RECODE_SESSIONERR)
		return
	}

	o := orm.NewOrm()
	user := models.User{Id: user_id.(int)}
	err := o.Read(&user)

	if err == orm.ErrNoRows {
		beego.Info("user 信息数据库查询不到")
		resp["errno"] = models.RECODE_USERERR
		resp["errmsg"] = models.RecodeText(models.RECODE_USERERR)
		return
	}
	//resp["data"] = user
	resData := make(map[string]interface{})
	/*
			"user_id": 1,
			"name": "Aceld",
		    "password": "123123",
			"mobile": "110",
		  	"real_name": "刘丹冰",
			"id_card": "210112244556677",
			"avatar_url": "http://101.200.170.171:9998/group1/M00/00/00/Zciqq1n7It2ANn1dAADexS5wJKs808.png"
	*/
	resData["user_id"] = user.Id
	resData["name"] = user.Name
	resData["password"] = user.Password_hash
	resData["mobile"] = user.Mobile
	resData["real_name"] = user.Real_name
	resData["id_card"] = user.Id_card
	resData["avatar_url"] = user.Avatar_url

	resp["data"] = resData
}
