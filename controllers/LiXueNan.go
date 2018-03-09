package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"lovehome/models"
	_ "path"
)

func (this *UserController) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

// 李雪楠_0_修改名字
func (this *UserController) PutName() {

	beego.Info("==========/api/v1.0/Auth post succ!!!=========")

	//返回给前端的map结构体
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)

	var regRequestMap = make(map[string]interface{})

	json.Unmarshal(this.Ctx.Input.RequestBody, &regRequestMap)

	user_id := this.GetSession("user_id")

	beego.Info("name = ", regRequestMap["name"])

	//2.将数据存入mysql数据库 user

	user := models.User{}
	user.Name = regRequestMap["name"].(string)
	user.Id = user_id.(int)

	o := orm.NewOrm()

	_, err := o.Update(&user, "name")
	if err != nil {
		beego.Info("insert  error =", err)
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	//3.将当前的用户信息存储到session中
	this.SetSession("name", user.Name)

	return

}

//李雪楠_1_实名验证请求
func (this *UserController) AuthPost() {
	beego.Info("==========/api/v1.0/Auth post succ!!!=========")

	//返回给前端的map结构体
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)

	var regRequestMap = make(map[string]interface{})

	json.Unmarshal(this.Ctx.Input.RequestBody, &regRequestMap)

	user_id := this.GetSession("user_id")
	id_card := regRequestMap["id_card"]
	real_name := regRequestMap["real_name"]

	beego.Info("id_card = ", regRequestMap["id_card"])
	beego.Info("id_card = ", regRequestMap["real_name"])

	//2.将数据存入mysql数据库 ser

	user := models.User{}
	user.Id_card = id_card.(string)
	user.Real_name = real_name.(string)
	user.Id = user_id.(int)

	o := orm.NewOrm()

	_, err := o.Update(&user, "id_card", "real_name")
	if err != nil {
		beego.Info("insert  error =", err)
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	return
}

//李雪楠_2_实名验证请求
func (this *UserController) AuthGet() {

	beego.Info("==========/api/v1.0/name post succ!!!=========")

	//返回给前端的map结构体
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)

	user_id := this.GetSession("user_id")

	//2.将数据存入mysql数据库 user

	user := models.User{}

	user.Id = user_id.(int)

	o := orm.NewOrm()

	err := o.Read(&user)
	if err != nil {
		beego.Info("insert  error =", err)
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	resp["data"] = user

	return

}
