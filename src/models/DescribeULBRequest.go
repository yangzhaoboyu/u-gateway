package models

//获取ULB请求参数
type DescribeULBRequest struct {
	Region    string `required:"true" description:"区域" json:"region"`
	ProjectId string `required:"true" description:"项目编号" json:"projectId"`
	Limit     int    `description:"数据分页值，默认为20" json:"limit"`
	Offset    int    `description:"数据偏移量，默认为0" json:"offset"`
}
