package servers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	. "github.com/woodylan/go-websocket/config"
	"github.com/woodylan/go-websocket/define/retcode"
	. "github.com/woodylan/go-websocket/models"
	"github.com/woodylan/go-websocket/pkg/setting"
	"github.com/woodylan/go-websocket/tools/util"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//channel通道
var ToClientChan chan clientInfo

//channel通道结构体
type clientInfo struct {
	ClientId   string
	SendUserId string
	MessageId  string
	Code       int
	Msg        string
	Data       interface{} //*string
}

type RetData struct {
	ClientId   string      `json:"clientId"`
	MessageId  string      `json:"messageId"`
	SendUserId string      `json:"sendUserId"`
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
}

// 心跳间隔
var heartbeatInterval = 60 * time.Second

func init() {
	ToClientChan = make(chan clientInfo, 1000)
}

var Manager = NewClientManager() // 管理者

func StartWebSocket() {
	websocketHandler := &Controller{}
	log.Println("进入websocket")
	http.HandleFunc("/ws", websocketHandler.Run)

	go Manager.Start()
}

//发送信息到指定客户端
func SendMessage2Client(clientId string, sendUserId string, code int, msg string, data interface{}) (messageId string) {
	messageId = util.GenUUID()
	if util.IsCluster() {
		addr, _, _, isLocal, err := util.GetAddrInfoAndIsLocal(clientId)
		if err != nil {
			log.Errorf("%s", err)
			return
		}

		//如果是本机则发送到本机
		if isLocal {
			SendMessage2LocalClient(messageId, clientId, sendUserId, code, msg, data)
		} else {
			//发送到指定机器
			SendRpc2Client(addr, messageId, sendUserId, clientId, code, msg, data)
		}
	} else {
		//如果是单机服务，则只发送到本机
		SendMessage2LocalClient(messageId, clientId, sendUserId, code, msg, data)
	}

	return
}

//关闭客户端
func CloseClient(clientId, systemId string) {
	if util.IsCluster() {
		addr, _, _, isLocal, err := util.GetAddrInfoAndIsLocal(clientId)
		if err != nil {
			log.Errorf("%s", err)
			return
		}

		//如果是本机则发送到本机
		if isLocal {
			CloseLocalClient(clientId, systemId)
		} else {
			//发送到指定机器
			CloseRpcClient(addr, clientId, systemId)
		}
	} else {
		//如果是单机服务，则只发送到本机
		CloseLocalClient(clientId, systemId)
	}

	return
}

//添加客户端到系统
func AddClientMap(clientSocket *Client) {
	Manager.EventConnect(clientSocket)
}

//添加客户端到分组
func AddClient2Group(systemId string, groupName string, clientId string, userId string, extend string) {

	if util.IsCluster() {
		//判断key是否存在
		addr, _, _, isLocal, err := util.GetAddrInfoAndIsLocal(clientId)
		if err != nil {
			log.Errorf("%s", err)
			return
		}

		if isLocal {
			if client, err := Manager.GetByClientId(clientId); err == nil {
				//添加到本地
				Manager.AddClient2LocalGroup(groupName, client, userId, extend)
			} else {
				log.Error(err)
			}
		} else {
			//发送到指定的机器
			SendRpcBindGroup(addr, systemId, groupName, clientId, userId, extend)
		}
	} else {
		//log.Println("这一步有问题吗？")
		if client, err := Manager.GetByClientId(clientId); err == nil {
			//如果是单机，就直接添加到本地group了
			//log.Printf("添加到本地机器列表group...%v\n", client)
			Manager.AddClient2LocalGroup(groupName, client, userId, extend)
		} else {
			log.Printf("Manager.GetByClientId 出什么问题了？%v\n", err)
		}
	}
}

//发送信息到指定分组
func SendMessage2Group(systemId, sendUserId, groupName string, code int, msg string, data interface{}) (messageId string) {
	messageId = util.GenUUID()
	if util.IsCluster() {
		//发送分组消息给指定广播
		go SendGroupBroadcast(systemId, messageId, sendUserId, groupName, code, msg, data)
	} else {
		//如果是单机服务，则只发送到本机
		Manager.SendMessage2LocalGroup(systemId, messageId, sendUserId, groupName, code, msg, data)
	}
	return
}

//发送信息到指定系统
func SendMessage2System(systemId, sendUserId string, code int, msg string, data string) {
	messageId := util.GenUUID()
	if util.IsCluster() {
		//发送到系统广播
		SendSystemBroadcast(systemId, messageId, sendUserId, code, msg, &data)
	} else {
		//如果是单机服务，则只发送到本机
		Manager.SendMessage2LocalSystem(systemId, messageId, sendUserId, code, msg, &data)
	}
}

//获取分组列表
func GetOnlineList(systemId *string, groupName *string) map[string]interface{} {
	var clientList []string
	if util.IsCluster() {
		//发送到系统广播
		clientList = GetOnlineListBroadcast(systemId, groupName)
	} else {
		//如果是单机服务，则只发送到本机
		retList := Manager.GetGroupClientList(util.GenGroupKey(*systemId, *groupName))
		clientList = append(clientList, retList...)
	}

	return map[string]interface{}{
		"count": len(clientList),
		"list":  clientList,
	}
}

//获取分组列表
func CheckOnline(groupName *string, clientId string) bool {

	if util.IsCluster() {
		//发送到系统广播
		//clientList = GetOnlineListBroadcast(systemId, groupName)
	} else {
		//如果是单机服务，则只发送到本机
		retList := Manager.GetGroupClientList(util.GenGroupKey(setting.CommonSetting.SystemId, *groupName))
		//log.Printf("所有用户%#v", retList)
		for _, item := range retList {
			if item == clientId {
				return true
			}
		}

	}
	return false
}

//通过本服务器发送信息
func SendMessage2LocalClient(messageId, clientId string, sendUserId string, code int, msg string, data interface{}) {
	log.WithFields(log.Fields{
		"host":     setting.GlobalSetting.LocalHost,
		"port":     setting.CommonSetting.HttpPort,
		"clientId": clientId,
	}).Info("发送到通道")
	ToClientChan <- clientInfo{ClientId: clientId, MessageId: messageId, SendUserId: sendUserId, Code: code, Msg: msg, Data: data}
	return
}

//发送关闭信号
func CloseLocalClient(clientId, systemId string) {
	if conn, err := Manager.GetByClientId(clientId); err == nil && conn != nil {
		if conn.SystemId != systemId {
			return
		}
		Manager.DisConnect <- conn
		log.WithFields(log.Fields{
			"host":     setting.GlobalSetting.LocalHost,
			"port":     setting.CommonSetting.HttpPort,
			"clientId": clientId,
		}).Info("主动踢掉客户端")
	}
	return
}

//监听并发送给客户端信息
func WriteMessage() {
	for {
		clientInfo := <-ToClientChan
		log.WithFields(log.Fields{
			"host":       setting.GlobalSetting.LocalHost,
			"port":       setting.CommonSetting.HttpPort,
			"clientId":   clientInfo.ClientId,
			"messageId":  clientInfo.MessageId,
			"sendUserId": clientInfo.SendUserId,
			"code":       clientInfo.Code,
			"msg":        clientInfo.Msg,
			//"data":       clientInfo.Data,
		}).Info("发送到本机")
		if conn, err := Manager.GetByClientId(clientInfo.ClientId); err == nil && conn != nil {
			if err := Render(conn.Socket, clientInfo.ClientId, clientInfo.MessageId, clientInfo.SendUserId, clientInfo.Code, clientInfo.Msg, clientInfo.Data); err != nil {
				Manager.DisConnect <- conn
				log.WithFields(log.Fields{
					"host":     setting.GlobalSetting.LocalHost,
					"port":     setting.CommonSetting.HttpPort,
					"clientId": clientInfo.ClientId,
					"msg":      clientInfo.Msg,
				}).Error("客户端异常离线：" + err.Error())
			}
		}
	}
}

func Render(conn *websocket.Conn, clientId string, messageId string, sendUserId string, code int, message string, data interface{}) error {
	return conn.WriteJSON(RetData{
		Code:       code,
		ClientId:   clientId,
		MessageId:  messageId,
		SendUserId: sendUserId,
		Msg:        message,
		Data:       data,
	})
}

//启动定时器进行心跳检测
func PingTimer() {
	go func() {
		ticker := time.NewTicker(heartbeatInterval)
		defer ticker.Stop()
		for {
			<-ticker.C
			//发送心跳
			for clientId, conn := range Manager.AllClient() {
				if err := conn.Socket.WriteControl(websocket.PingMessage, []byte("heartbeat"), time.Now().Add(time.Second)); err != nil {
					Manager.DisConnect <- conn
					log.Printf("发送心跳失败: %s 总连接数：%d", clientId, Manager.Count())
				} else {
					//log.Printf("发送心跳....%s", clientId)
				}
			}
		}

	}()
}
func RedisSend() {
	var ctx = context.Background()
	redisSubscribe := RedisClient.Subscribe(ctx, ViperConfig.Redis.Key)
	_, err := redisSubscribe.Receive(ctx)
	if err != nil {
		log.Printf("redis 订阅出错 %v\n", err)
		log.Fatal(err)
	}
	for msg := range redisSubscribe.Channel() {
		msg.Payload = strings.Trim(msg.Payload, "\"")
		log.Printf("redis读取数据：channel=%s", msg.Channel)
		if strings.TrimSpace(msg.Payload) == "" {
			log.Printf("空消息...")
			continue
		}
		str := []byte(msg.Payload)
		data := MsgType{}
		err := json.Unmarshal(str, &data)
		if err != nil {
			log.Printf("消息解析出错...%v", err)
			continue
		}
		onLine := Manager.Count()
		data.Sub.OnLine = onLine
		log.Printf("解析结果%v", data)
		SendUserId := strconv.Itoa(data.Msg.UserId)
		GroupName := strconv.Itoa(data.Msg.ChatroomId)
		SendMessage2Group(setting.CommonSetting.SystemId, SendUserId, GroupName, 200, "success", data)

	}

}

func SendListMsgToClient(manager *ClientManager, clientId string, userId string, groupName string) {
	data := GetList(groupName, clientId, 0, 1)
	data.Sub.OnLine = manager.Count()

	SendMessage2Client(clientId, userId, retcode.ONLINE_MESSAGE_CODE, "客户端上线", data)
}
func GetList(groupName string, clientId string, lastId int, page int) (data Lists) {
	msg := Msg{}
	list, _ := msg.GetList(lastId)
	timeUnix := time.Now().Unix()
	for k, v := range list {
		t1, _ := time.Parse(time.RFC3339, v.CreateTime)
		list[k].CreateTime = t1.Format("2006-01-02 15:04:05")
	}
	i := 0
	for k, v := range list {
		i++
		shanghaiZone, _ := time.LoadLocation("Asia/Shanghai")
		times, _ := time.ParseInLocation("2006-01-02 15:04:05", v.CreateTime, shanghaiZone)
		//log.Printf("数据库取出时间%s,原始时间%s,现在的时间%d,数据的时间%d\n", v.CreateTime, times, timeUnix, times.Unix())
		if i+1 >= len(list) {
			list[k].TimeMsg = TimeMsgNUll{}
		} else {
			temp := list[i+1]
			t_times, _ := time.ParseInLocation("2006-01-02 15:04:05", temp.CreateTime, shanghaiZone)
			if i == 1 && (times.Unix()-t_times.Unix()) >= 180 {

				//第一条时间大于大二条
				tmu := util.GetChatTimeStr(times.Unix())
				tm := TimeMsg{tmu}
				list[k].TimeMsg = tm
				timeUnix = t_times.Unix()
				i = 0

			} else if i >= 10 && (timeUnix-times.Unix()) < 180 {
				tmu := util.GetChatTimeStr(times.Unix())
				tm := TimeMsg{tmu}
				list[k].TimeMsg = tm
				timeUnix = times.Unix()
				i = 0
				//log.Println("时间消息---条数到了...")
			} else if i <= 10 && (timeUnix-times.Unix()) >= 180 {
				tmu := util.GetChatTimeStr(times.Unix())
				tm := TimeMsg{tmu}
				list[k].TimeMsg = tm
				//log.Println("时间消息---时间到了...")
				timeUnix = times.Unix()
				i = 0
			} else {
				list[k].TimeMsg = TimeMsgNUll{}
			}
		}

	}
	//data := Lists{}
	num := len(list)

	data.List = list
	tid, err := msg.GetTop()
	if err != nil {
		log.Printf(" 消息 id 出错 %v\n", err)
		//log.Fatal(err)
	}
	//manager := ClientManager{}
	//var manager = NewClientManager()
	onLine := Manager.Count()
	//on := GetOnlineList(&setting.CommonSetting.SystemId, &groupName)
	//	onLine, _ := strconv.Atoi(on["count"])

	ChatroomId, err := strconv.Atoi(groupName)
	subs := Subs{ChatroomId, 0, page, tid, onLine, 2, clientId}
	if num > 0 {
		var ctx = context.Background()
		tmp := list[num-1]
		RedisClient.HSet(ctx, clientId, "last_id", tmp.Id, "page", page)
		subs = Subs{ChatroomId, tmp.Id, page, tid, onLine, 2, clientId}
	} else {
		data.List = ListNull{}
	}

	data.Sub = subs
	return
}
