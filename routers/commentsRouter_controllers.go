package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["OneeSan/controllers:IndexController"] = append(beego.GlobalControllerRouter["OneeSan/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OneeSan/controllers:IndexController"] = append(beego.GlobalControllerRouter["OneeSan/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexAbout",
            Router: `/about`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OneeSan/controllers:IndexController"] = append(beego.GlobalControllerRouter["OneeSan/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexApi",
            Router: `/api`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OneeSan/controllers:IndexController"] = append(beego.GlobalControllerRouter["OneeSan/controllers:IndexController"],
        beego.ControllerComments{
            Method: "ApiAddIllust",
            Router: `/api/addillust`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OneeSan/controllers:IndexController"] = append(beego.GlobalControllerRouter["OneeSan/controllers:IndexController"],
        beego.ControllerComments{
            Method: "ApiDBCount",
            Router: `/api/dbcount`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OneeSan/controllers:IndexController"] = append(beego.GlobalControllerRouter["OneeSan/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexRobots",
            Router: `/robots.txt`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OneeSan/controllers:IndexController"] = append(beego.GlobalControllerRouter["OneeSan/controllers:IndexController"],
        beego.ControllerComments{
            Method: "IndexTest",
            Router: `/test`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
