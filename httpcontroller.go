package msgserver

import (
	"git.qianqiusoft.com/qianqiusoft/ssoclient"

	"github.com/astaxie/beego"
)

type HttpController struct {
	beego.Controller
	user *ssoclient.Token
}

func (this *HttpController) Prepare() {

	this.user = ssoclient.ValidateController(this.Ctx.ResponseWriter, this.Ctx.Request)
	if this.user == nil {
		return
	}
}
