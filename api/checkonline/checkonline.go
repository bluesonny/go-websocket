package checkonline

import (
	log "github.com/sirupsen/logrus"
	"github.com/woodylan/go-websocket/api"
	"github.com/woodylan/go-websocket/define/retcode"
	"github.com/woodylan/go-websocket/servers"
	"net/http"
)

type Controller struct {
}

type inputData struct {
	ChatroomId string      `json:"chatroom_id" validate:"required"`
	ClientId   string      `json:"client_id" validate:"required"`
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}

func (c *Controller) Run(w http.ResponseWriter, request *http.Request) {
	var inputData inputData
	//if err := json.NewDecoder(r.Body).Decode(&inputData); err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//err := request.ParseMultipartForm(1024)
	err := request.ParseForm()
	//log.Printf("Cannot parse form%#v", request.Form)
	if err != nil {
		log.Println("Cannot parse form")
	}
	ChatroomId := request.PostFormValue("chatroom_id")
	ClientId := request.PostFormValue("client_id")
	log.Printf("参数值ChatroomId：%v，ClientId：%v", ChatroomId, ClientId)
	if ClientId == "" || ChatroomId == "" {
		//errMap := map[string]string{"code": "1001", "msg": "参数不全"}
		api.Render(w, retcode.ONLINE_MESSAGE_CODE, "success", false)
		return
	}
	inputData.ClientId = ClientId
	inputData.ChatroomId = ChatroomId
	ret := servers.CheckOnline(&inputData.ChatroomId, inputData.ClientId)
	api.Render(w, retcode.SUCCESS, "success", ret)
	return
}
