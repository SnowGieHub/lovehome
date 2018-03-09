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
	} else {
		resp["data"] = user
	}
}
