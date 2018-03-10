package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"lovehome/models"
	"strconv"
)

type HouseController struct {
	beego.Controller
}

func (this *HouseController) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *HouseController) GetHouseDetail() {
	beego.Info("GetHouseDetail .....")
	house_id := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(house_id)
	myHouse := models.House{Id: id}
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer this.RetData(resp)

	o := orm.NewOrm()
	//	qs := o.QueryTable("house")
	//	err := qs.Filter("Id", house_id).RelatedSel().One(&myHouse)
	err := o.Read(&myHouse)
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
	}

	o.LoadRelated(&myHouse, "User")
	o.LoadRelated(&myHouse, "Area")
	o.LoadRelated(&myHouse, "Facilities")
	o.LoadRelated(&myHouse, "Images")
	o.LoadRelated(&myHouse, "Orders")

	beego.Info(myHouse)
	resp["errno"] = 0
	resp["errmsg"] = "成功"
	myData := make(map[string]interface{})
	houseData := make(map[string]interface{})

	//myData["house"] = myHouse
	myData["user_id"] = this.GetSession("user_id")
	resp["data"] = myData

	houseData["acreage"] = myHouse.Acreage
	houseData["address"] = myHouse.Address
	houseData["beds"] = myHouse.Beds
	houseData["capacity"] = myHouse.Capacity
	houseData["deposit"] = myHouse.Deposit
	houseData["hid"] = house_id
	houseData["max_days"] = myHouse.Max_days
	houseData["min_days"] = myHouse.Min_days
	houseData["price"] = myHouse.Price
	houseData["room_count"] = myHouse.Room_count
	houseData["title"] = myHouse.Title
	houseData["unit"] = myHouse.Unit
	houseData["user_avatar"] = myHouse.User.Avatar_url
	houseData["user_id"] = this.GetSession("user_id")
	houseData["user_name"] = myHouse.User.Name

	beego.Info(houseData)

	comments := make([]string, len(myHouse.Orders))
	//for houseData["comments"] = myHouse.
	for _, order := range myHouse.Orders {
		beego.Info("=================", order)
		comments = append(comments, order.Comment)
	}
	beego.Info("===========comments:", comments)
	beego.Info("len() == ", len(myHouse.Orders))
	houseData["comments"] = comments

	facilities := make([]int, len(myHouse.Facilities))
	//for houseData["facilities"] = myHouse.
	for _, facility := range myHouse.Facilities {
		beego.Info("=================", facility)
		facilities = append(facilities, facility.Id)
	}
	beego.Info("===========facilities:", facilities)
	beego.Info("len() == ", len(myHouse.Facilities))
	houseData["facilities"] = facilities

	imgs := make([]string, len(myHouse.Images))
	//for houseData["img_urls"] = myHouse.
	for _, img := range myHouse.Images {
		imgs = append(imgs, "http://192.168.5.129:9988/"+img.Url)
	}
	houseData["img_urls"] = imgs

	myData["house"] = houseData
	resp["data"] = myData

	return
}
