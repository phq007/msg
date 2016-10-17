package msgserver

import (
	"git.qianqiusoft.com/qianqiusoft/ssoclient"

	"github.com/astaxie/beego"
)

type WebsocketController struct {
	beego.Controller
	user *ssoclient.Token
}

func (this *WebsocketController) Prepare() {

	this.user = ssoclient.ValidateController(this.Ctx.ResponseWriter, this.Ctx.Request)
	if this.user == nil {
		return
	}
}

func (this *WebsocketController) Join() {
	udevice := this.GetString("devicetoken")
	uip := this.Ctx.Input.IP()
	// Upgrade from http request to WebSocket.
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	c := NewConnection(ws, make(chan []byte, 256))

	//Register this connection into the (global var) hub
	h := GetHub(id)
	h.Register(c)

	//If this deferred function gets called, it implies that
	// writer and reader already exited
	defer func() {
		h.Unregister(c)
	}()
	go c.Writer()
	c.Reader(h)
}
