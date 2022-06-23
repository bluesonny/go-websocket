package servers

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/woodylan/go-websocket/define/retcode"
	"github.com/woodylan/go-websocket/pkg/setting"
	"github.com/woodylan/go-websocket/tools/util"
	"net/http"
)

const (
	// 最大的消息大小
	maxMessageSize = 8192
)

type Controller struct {
}

type renderData struct {
	ClientId string `json:"clientId"`
}

func (c *Controller) Run(w http.ResponseWriter, r *http.Request) {
	log.Println("websocket 方法开始运行...")
	conn, err := (&websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// 允许所有CORS跨域请求
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("upgrade error: %v", err)
		http.NotFound(w, r)
		return
	}

	//设置读取消息大小上线
	conn.SetReadLimit(maxMessageSize)

	//解析参数
	systemId := setting.CommonSetting.SystemId //r.FormValue("systemId")
	userId := r.FormValue("user_id")
	chat_type := r.FormValue("chat_type")
	if len(userId) == 0 {
		_ = Render(conn, "", "", "", retcode.SYSTEM_ID_ERROR, "用户ID不能为空", []string{})
		_ = conn.Close()
		return
	}
	if len(chat_type) == 0 {
		chat_type = "chat" //bullet_chat  弹幕
	}

	clientId := util.GenClientId()

	clientSocket := NewClient(clientId, systemId, conn)

	Manager.AddClient2SystemClient(systemId, clientSocket)

	//读取客户端消息--监听客户端
	clientSocket.Read()

	//if err = api.ConnRender(conn, renderData{ClientId: clientId}); err != nil {
	//	_ = conn.Close()
	//	return
	//}

	// 用户连接事件
	//Manager.Connect <- clientSocket
	AddClientMap(clientSocket)
	//连接成功加入群组
	//chatroomId := r.FormValue("chatroomId")
	//if len(chatroomId) == 0 {
	chatroomId := setting.CommonSetting.ChatroomId //默认群组
	log.Println("获取默认群组...")
	//}
	if chat_type == "bullet_chat" {
		chatroomId = setting.CommonSetting.BulletChat //默认群组
		log.Println("---自动加入默认弹幕群组---")
	}

	AddClient2Group(systemId, chatroomId, clientId, userId, chat_type)
}
