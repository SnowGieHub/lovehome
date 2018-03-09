package controllers

import (
	_ "encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"lovehome/models"
	"path"
)

type UpImageControllers struct {
	beego.Controller
}

func (this *UpImageControllers) RetData(resp interface{}) {
	this.Data["json"] = resp
	this.ServeJSON()
}

func (this *UpImageControllers) UpHouseImage() {
	resp := make(map[string]interface{})
	resp["errno"] = models.RECODE_OK
	resp["errmsg"] = models.RECODE_OK

	defer this.RetData(resp)

	file, header, err := this.GetFile("images")
	if err != nil {
		resp["errno"] = models.RECODE_SERVERERR
		resp["errmsg"] = models.RecodeText(models.RECODE_SERVERERR)
		return
	}

	filebuffer := make([]byte, header.Size)
	if _, err := file.Read(filebuffer); err != nil {

		resp["errno"] = models.RECODE_IOERR
		resp["errmsg"] = models.RecodeText(models.RECODE_IOERR)
		return
	}

	suffix := path.Ext(header.Filename)

	groupName, fileId, err := models.FDFSUploadByBuffer(filebuffer, suffix[1:])
	if err != nil {

		resp["errno"] = models.RECODE_IOERR
		resp["errmsg"] = models.RecodeText(models.RECODE_IOERR)
		return
	}
	beego.Info("fdfs upload succ groupname = ", groupName, " fileid = ", fileId)

	house_id := this.GetSession("house_id")
	house := models.House{Id: house_id.(int)}

	o := orm.NewOrm()
	err = o.Read(&house)
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RECODE_DBERR
		return
	}
	image := models.HouseImage{}

	image_url := "http://192.168.191.140:8080/" + fileId
	image.House = &house
	image.Url = image_url
	url_map := make(map[string]interface{})
	_, err = o.Insert(&image)
	if err != nil {
		resp["errno"] = models.RECODE_DBERR
		resp["errmsg"] = models.RECODE_DBERR
		return
	}
	url_map["avatar_url"] = image_url

	resp["data"] = url_map

	return
}
