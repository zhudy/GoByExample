package router

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"../controller"
)
func init(){
	logs.Debug("enter route init")
	beego.Router("/seckill", &controller.SkillController{}, "*:SecKill")  //匹配，调用
	beego.Router("/secinfo", &controller.SkillController{}, "*:SecInfo")
}