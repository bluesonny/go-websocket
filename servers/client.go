package servers

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"time"
)

type Client struct {
	ClientId    string          // 标识ID
	SystemId    string          // 系统ID
	Socket      *websocket.Conn // 用户连接
	ConnectTime uint64          // 首次连接时间
	IsDeleted   bool            // 是否删除或下线
	UserId      string          // 业务端标识用户ID
	Extend      string          // 扩展字段，用户可以自定义
	GroupList   []string
}

type SendData struct {
	Code int
	Msg  string
	Data *interface{}
}

func NewClient(clientId string, systemId string, socket *websocket.Conn) *Client {
	return &Client{
		ClientId:    clientId,
		SystemId:    systemId,
		Socket:      socket,
		ConnectTime: uint64(time.Now().Unix()),
		IsDeleted:   false,
	}
}

func (c *Client) Read() {
	go func() {
		for {
			messageType, _, err := c.Socket.ReadMessage()
			log.Printf("循环读取消息出错了,连接断开了?messageType:%v;err:%v,", messageType, err)
			if err != nil {
				if messageType == -1 { //&& websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure, websocket.CloseNoStatusReceived, websocket.CloseAbnormalClosure) {
					Manager.DisConnect <- c
					return
				} //else if messageType != websocket.PingMessage || messageType != websocket.PongMessage {
				//log.Printf("循环读取消息出错了--心跳信息:messageType:%v,%v", messageType, err)
				//return
				//}
			}
		}
	}()
}
