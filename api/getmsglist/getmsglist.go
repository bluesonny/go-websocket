package getmsglist

import (
	log "github.com/sirupsen/logrus"
	"github.com/woodylan/go-websocket/api"
	"github.com/woodylan/go-websocket/define/retcode"
	"github.com/woodylan/go-websocket/models"
	"github.com/woodylan/go-websocket/servers"
	"net/http"
	"strconv"
)

type Controller struct {
}

type inputData struct {
	ChatroomId string `json:"chatroom_id" validate:"required"`
	ClientId   string `json:"client_id" validate:"required"`
	//UserId     int         `json:"user_id" validate:"required"`
	LastId int         `json:"last_id"  validate:"required"`
	Page   int         `json:"page"`
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
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
	LastId := request.PostFormValue("last_id")
	Page := request.PostFormValue("page")
	log.Printf("请求消息列表接口参数---ChatroomId：%v，ClientId:%s,LastId：%v,Page:%v", ChatroomId, ClientId, LastId, Page)
	lastId, err := strconv.Atoi(LastId)
	page, err := strconv.Atoi(Page)
	if err != nil {
		log.Printf("出什么错误了%v", err)
		errMap := map[string]string{"code": "1001", "msg": "参数不全"}
		api.Render(w, retcode.ONLINE_MESSAGE_CODE, "success", errMap)
		return
	}
	if page == 0 {
		page = 1
	}
	page = page + 1
	errMap := map[string]string{"code": "1002", "msg": "无法请求更多数据"}
	if page > 20 {

		api.Render(w, retcode.ONLINE_MESSAGE_CODE, "success", errMap)
		return
	}
	inputData.ChatroomId = ChatroomId
	inputData.ClientId = ClientId
	inputData.LastId = lastId
	inputData.Page = page

	//err = api.Validate(inputData)
	if err != nil {
		errMap := map[string]string{"code": "1002", "msg": err.Error()}
		api.Render(w, retcode.ONLINE_MESSAGE_CODE, "success", errMap)
		return
	}
	ret := servers.GetList(inputData.ChatroomId, inputData.ClientId, inputData.LastId, page)
	//log.Printf("返回：%#v", ret)
	if _, ok := ret.List.(models.ListNull); ok {
		//log.Printf("%v", n)
		api.Render(w, retcode.ONLINE_MESSAGE_CODE, "success", errMap)
		return
	}
	api.Render(w, retcode.SUCCESS, "success", ret)
	return
}
