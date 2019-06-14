// @APIVersion 1.0.0
// @Title ucloud网关
// @Description 本API提供ucloud连接网关服务 现已支持ULB
// @TermsOfServiceUrl https://docs.ucloud.cn/api/ulb-api/index
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"u-gateway/controllers"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/ulbs",
			beego.NSInclude(
				&controllers.UlBController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
