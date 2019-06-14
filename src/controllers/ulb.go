package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/ucloud/ucloud-sdk-go/services/ulb"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/log"
	"u-gateway/models"
)

var (
	UConf       ucloud.Config
	UCredential auth.Credential
)

func init() {
	UConf = ucloud.NewConfig()
	UConf.LogLevel = log.DebugLevel
	UCredential = auth.NewCredential()
	UCredential.PublicKey = beego.AppConfig.String("PublicKey")
	UCredential.PrivateKey = beego.AppConfig.String("PrivateKey")
}

//ULB相关
type UlBController struct {
	beego.Controller
}

// @Title 获取ULB列表
// @Description 获取ULB列表
// @Param	body		body 	models.DescribeULBRequest	true		"请求参数"
// @Success 200 {object} github.com.ucloud.ucloud-sdk-go.services.ulb.DescribeULBResponse  "内嵌包swagger无法显示,请参考ucloud文档"
// @router /query [post]
func (this *UlBController) Post() {
	var req models.DescribeULBRequest
	json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if req.Region == "" || req.Region == "string" {
		this.Data["json"] = map[string]string{"error": "区域不可为空"}
		this.ServeJSON()
		this.StopRun()
	}
	if req.ProjectId == "" || req.ProjectId == "string" {
		this.Data["json"] = map[string]string{"error": "项目编号不可为空"}
		this.ServeJSON()
		this.StopRun()
	}
	ulbClient := ulb.NewClient(&UConf, &UCredential)
	request := ulbClient.NewDescribeULBRequest()
	request.Region = ucloud.String(req.Region)
	request.ProjectId = ucloud.String(req.ProjectId)
	if req.Limit <= 0 {
		req.Limit = 20
	}
	request.Limit = ucloud.Int(req.Limit)
	request.Offset = ucloud.Int(req.Offset)
	response, err := ulbClient.DescribeULB(request)
	if err != nil {
		this.Data["json"] = map[string]string{"error": err.Error()}

	} else {
		this.Data["json"] = response
	}
	this.ServeJSON()
}

// @Title 更新ULB下UServer下的服务节点状态
// @Description 更新ULB下UServer下的服务节点状态
// @Param	body		body 	models.UpdateBackendAttributeRequest	true		"请求参数"
// @Success 200 {object} models.BaseResponse
// @router /update [post]
func (this *UlBController) Update() {
	var req models.UpdateBackendAttributeRequest
	json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	if req.Region == "" || req.Region == "string" {
		this.Data["json"] = map[string]string{"error": "区域不可为空"}
		this.ServeJSON()
		this.StopRun()
	}
	if req.ProjectId == "" || req.ProjectId == "string" {
		this.Data["json"] = map[string]string{"error": "项目编号不可为空"}
		this.ServeJSON()
		this.StopRun()
	}
	if req.ULBId == "" || req.ULBId == "string" {
		this.Data["json"] = map[string]string{"error": "ULB编号不可为空"}
		this.ServeJSON()
		this.StopRun()
	}
	if req.BackendId == "" || req.BackendId == "string" {
		this.Data["json"] = map[string]string{"error": "归属ULB下的服务节点不可为空"}
		this.ServeJSON()
		this.StopRun()
	}
	ulbClient := ulb.NewClient(&UConf, &UCredential)
	request := ulbClient.NewUpdateBackendAttributeRequest()
	request.Region = ucloud.String(req.Region)
	request.ProjectId = ucloud.String(req.ProjectId)
	request.ULBId = ucloud.String(req.ULBId)
	request.BackendId = ucloud.String(req.BackendId)
	request.Enabled = ucloud.Int(req.Enabled)
	response, err := ulbClient.UpdateBackendAttribute(request)
	if err != nil {
		this.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		resp := models.BaseResponse{}
		resp.Action = response.Action
		resp.Message = response.Message
		resp.RetCode = response.RetCode
		this.Data["json"] = resp
	}
	this.ServeJSON()
}
