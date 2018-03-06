package controllers

import (
	"github.com/astaxie/beego"
	"lovehome/models"
)

type SessionController struct {
	beego.Controller
}

//将封装好的返回结构 变成json返回给前段
func (this *SessionController) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *SessionController) GetSessionName() {

	beego.Info("=============/api/v1.0/session get Session succ ==============")

	resq := make(map[string]interface{})

	resq["errno"] = models.RECODE_SESSIONERR
	resq["errmsg"] = models.RecodeText(models.RECODE_SERVERERR)

	defer this.RetData(resq)

	name_map := make(map[string]interface{})

	name := this.GetSession("name")

	if name != nil {

		resq["errno"] = models.RECODE_OK
		resq["errmsg"] = models.RecodeText(models.RECODE_OK)
		name_map["name"] = name.(string)
		resq["data"] = name_map

	}

	return

}

// /api/v1.0/session
func (this *SessionController) DelSessionName() {
	beego.Info("========== /api/v1.0/session get Session succ ======")

	resp := make(map[string]interface{})

	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)

	this.DelSession("name")
	this.DelSession("user_id")
	this.DelSession("mobile")

	return
}
