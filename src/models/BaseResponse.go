package models

//基础响应信息
type BaseResponse struct {
	Action  string `json:"action" description:"操作"`
	RetCode int    `json:"retCode" description:"响应码(0:成功)"`
	Message string `json:"message" description:"输出信息"`
}
