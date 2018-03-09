package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"lovehome/models"
	"strconv"
)

/*
	路由路径
	beego.Router("/api/v1.0/houses", &controllers.HousesController{}, "post:ReleaseHouseInfo")
	beego.Router("/api/v1.0/houses", &controllers.HousesController{}, "post:ReleaseHouseInfo")
	beego.Router("/api/v1.0/houses/:id/images", &controllers.UpImageControllers{} "post:UpHouseImage")
*/
type HousesController struct {
	beego.Controller
}

func (this *HousesController) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *HousesController) ReleaseHouseInfo() {
	beego.Info("==========/api/v1.0/house!!!=========")

	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)

	defer this.RetData(resp)

	user_id := this.GetSession("user_id")
	user := models.User{Id: user_id.(int)}

	o := orm.NewOrm()
	o.Read(&user)

	var regRequestMap = make(map[string]interface{})

	json.Unmarshal(this.Ctx.Input.RequestBody, &regRequestMap)

	house := models.House{}

	array := regRequestMap["facility"].([]interface{})
	lenth := len(array)
	facilites_arr := make([]*models.Facility, lenth)
	for i := 0; i < lenth; i++ {
		facilites_arr[i] = new(models.Facility)
		facilites_arr[i].Id, _ = strconv.Atoi(array[i].(string))
		o.Read(facilites_arr[i])
	}

	house.Title = regRequestMap["title"].(string)
	house.Price, _ = strconv.Atoi(regRequestMap["price"].(string))
	house.Address = regRequestMap["address"].(string)
	house.Room_count, _ = strconv.Atoi(regRequestMap["room_count"].(string))
	house.Acreage, _ = strconv.Atoi(regRequestMap["acreage"].(string))
	house.Unit = regRequestMap["unit"].(string)
	house.Capacity, _ = strconv.Atoi(regRequestMap["capacity"].(string))
	house.Beds = regRequestMap["beds"].(string)
	house.Deposit, _ = strconv.Atoi(regRequestMap["deposit"].(string))
	house.Min_days, _ = strconv.Atoi(regRequestMap["min_days"].(string))
	house.Max_days, _ = strconv.Atoi(regRequestMap["max_days"].(string))
	house.User = &user
	area_id, _ := strconv.Atoi(regRequestMap["area_id"].(string))
	area := models.Area{Id: area_id}
	o.Read(&area)
	house.Area = &area
	house.Facilities = facilites_arr
	id, err := o.Insert(&house)
	if err != nil {
		beego.Info("insert error = ", err)
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}

	m2m := o.QueryM2M(&house, "facility_houses")

	o.Insert(m2m)
	/*
		_, err = o.InsertMulti(lenth, facilites_arr)
		if err != nil {
			beego.Info("insert error = ", err)
			resp["errno"] = models.RECODE_DBERR
			resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
			return
		}
		i*/
	house_id := make(map[string]interface{})
	house_id["house_id"] = id
	resp["data"] = house_id
	//4 给前段返回注册成功还是失败的结果
	return

}
func (this *HousesController) GetHouseInfo() {
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

	var posts []*models.House
	num, err := o.QueryTable("house").Filter("user_id", user_id.(int)).RelatedSel().All(&posts)
	if err == nil {
		fmt.Printf("%d posts read\n", num)
	}
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)
		return
	}
	resp["data"] = posts
	return
}
