package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["u-gateway/controllers:UlBController"] = append(beego.GlobalControllerRouter["u-gateway/controllers:UlBController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/query`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["u-gateway/controllers:UlBController"] = append(beego.GlobalControllerRouter["u-gateway/controllers:UlBController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/update`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
