package models

//更新服务节点请求参数
type UpdateBackendAttributeRequest struct {
	Region    string `required:"true" description:"区域" json:"region"`
	ProjectId string `required:"true" description:"项目编号" json:"projectId"`
	ULBId     string `required:"true" description:"ULB编号" json:"bLBId"`
	BackendId string `required:"true" description:"归属ULB下服务节点编号" json:"backendId"`
	Enabled   int    `required:"true" description:"状态(1:开启,0:关闭)" json:"enabled"`
}
