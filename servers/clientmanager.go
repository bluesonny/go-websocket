package servers

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/woodylan/go-websocket/define/retcode"
	. "github.com/woodylan/go-websocket/models"
	"github.com/woodylan/go-websocket/pkg/setting"
	"github.com/woodylan/go-websocket/tools/util"
	"strconv"
	"sync"
	"time"
)

// 连接管理
type ClientManager struct {
	ClientIdMap     map[string]*Client // 全部的连接
	ClientIdMapLock sync.RWMutex       // 读写锁

	Connect    chan *Client // 连接处理
	DisConnect chan *Client // 断开连接处理

	GroupLock sync.RWMutex
	Groups    map[string][]string

	SystemClientsLock sync.RWMutex
	SystemClients     map[string][]string
}

func NewClientManager() (clientManager *ClientManager) {
	clientManager = &ClientManager{
		ClientIdMap:   make(map[string]*Client),
		Connect:       make(chan *Client, 10000),
		DisConnect:    make(chan *Client, 10000),
		Groups:        make(map[string][]string, 100),
		SystemClients: make(map[string][]string, 100),
	}

	return
}

// 管道处理程序
func (manager *ClientManager) Start() {
	for {
		select {
		//case client := <-manager.Connect:
		// 建立连接事件
		//manager.EventConnect(client)
		case conn := <-manager.DisConnect:
			// 断开连接事件
			manager.EventDisconnect(conn)
		}
	}
}

// 建立连接事件
func (manager *ClientManager) EventConnect(client *Client) {
	manager.AddClient(client)

	log.WithFields(log.Fields{
		"host":     setting.GlobalSetting.LocalHost,
		"port":     setting.CommonSetting.HttpPort,
		"clientId": client.ClientId,
		"counts":   Manager.Count(),
	}).Info("客户端已连接")
}

// 断开连接时间
func (manager *ClientManager) EventDisconnect(client *Client) {
	//关闭连接
	_ = client.Socket.Close()
	manager.DelClient(client)

	//mJson, _ := json.Marshal(map[string]string{
	//	"clientId": client.ClientId,
	//	"userId":   client.UserId,
	//	"extend":   client.Extend,
	//})
	//data := string(mJson)
	//sendUserId := ""

	//发送下线通知
	if len(client.GroupList) > 0 {
		//for _, groupName := range client.GroupList {
		//SendMessage2Group(client.SystemId, sendUserId, groupName, retcode.OFFLINE_MESSAGE_CODE, "客户端下线", &data)
		//}
	}

	log.WithFields(log.Fields{
		"host":     setting.GlobalSetting.LocalHost,
		"port":     setting.CommonSetting.HttpPort,
		"clientId": client.ClientId,
		"counts":   Manager.Count(),
		"seconds":  uint64(time.Now().Unix()) - client.ConnectTime,
	}).Info("客户端已断开")

	//标记销毁
	client.IsDeleted = true
	client = nil
}

// 添加客户端
func (manager *ClientManager) AddClient(client *Client) {
	manager.ClientIdMapLock.Lock()
	defer manager.ClientIdMapLock.Unlock()
	//log.Printf("添加客户连接前%v,ClientId--%v", manager.ClientIdMap, client.ClientId)
	manager.ClientIdMap[client.ClientId] = client
	//log.Printf("添加客户连接后%v,ClientId--%v", manager.ClientIdMap, client.ClientId)
}

// 获取所有的客户端
func (manager *ClientManager) AllClient() map[string]*Client {
	manager.ClientIdMapLock.RLock()
	all := make(map[string]*Client, len(manager.ClientIdMap))
	for k, v := range manager.ClientIdMap {
		all[k] = v
	}

	manager.ClientIdMapLock.RUnlock()
	return all
}

// 客户端数量
func (manager *ClientManager) Count() int {
	manager.ClientIdMapLock.RLock()
	defer manager.ClientIdMapLock.RUnlock()
	return len(manager.ClientIdMap)
}

//分组客户端量
func (manager *ClientManager) GroupCount(groupKey string) int {
	manager.GroupLock.RLock()
	defer manager.GroupLock.RUnlock()
	return len(manager.Groups[groupKey])

}

// 删除客户端
func (manager *ClientManager) DelClient(client *Client) {
	manager.delClientIdMap(client.ClientId)

	//删除所在的分组
	if len(client.GroupList) > 0 {
		for _, groupName := range client.GroupList {
			manager.delGroupClient(util.GenGroupKey(client.SystemId, groupName), client.ClientId)
		}
	}

	// 删除系统里的客户端
	manager.delSystemClient(client)
}

// 删除clientIdMap
func (manager *ClientManager) delClientIdMap(clientId string) {
	manager.ClientIdMapLock.Lock()
	defer manager.ClientIdMapLock.Unlock()

	delete(manager.ClientIdMap, clientId)
}

// 通过clientId获取
func (manager *ClientManager) GetByClientId(clientId string) (*Client, error) {
	manager.ClientIdMapLock.RLock()
	defer manager.ClientIdMapLock.RUnlock()
	//log.Printf("所有连接--%v,目前clientId--%v", manager.ClientIdMap, clientId)

	if client, ok := manager.ClientIdMap[clientId]; !ok {
		log.Printf("通过clientId获取 出问题了么？%v\n", ok)
		return nil, errors.New("客户端不存在")
	} else {
		return client, nil
	}
}

// 发送到本机分组
func (manager *ClientManager) SendMessage2LocalGroup(systemId, messageId, sendUserId, groupName string, code int, msg string, data interface{}) {
	if len(groupName) > 0 {
		clientIds := manager.GetGroupClientList(util.GenGroupKey(systemId, groupName))
		if len(clientIds) > 0 {
			for _, clientId := range clientIds {
				if _, err := Manager.GetByClientId(clientId); err == nil {

					if groupName == setting.CommonSetting.BulletChat {
						bt, _ := data.(BulletChat)
						bt.ClientId = clientId
						SendMessage2LocalClient(messageId, clientId, sendUserId, code, msg, bt)
						continue
					}

					//添加到本地
					op, _ := data.(MsgType)
					op.Sub.ClientId = clientId

					var ctx = context.Background()
					last_id, err := RedisClient.HGet(ctx, clientId, "last_id").Result()
					page, err := RedisClient.HGet(ctx, clientId, "page").Result()

					if err != nil {
						log.Printf("读取缓存有问题？%v,%s,--%s", err, last_id, page)
					} else {
						op.Sub.LastId, err = strconv.Atoi(last_id)
						op.Sub.Page, err = strconv.Atoi(page)
					}
					SendMessage2LocalClient(messageId, clientId, sendUserId, code, msg, op)
				} else {
					//删除分组
					manager.delGroupClient(util.GenGroupKey(systemId, groupName), clientId)
				}
			}
		}
	}
}

//发送给指定业务系统
func (manager *ClientManager) SendMessage2LocalSystem(systemId, messageId string, sendUserId string, code int, msg string, data *string) {
	if len(systemId) > 0 {
		clientIds := Manager.GetSystemClientList(systemId)
		if len(clientIds) > 0 {
			for _, clientId := range clientIds {
				SendMessage2LocalClient(messageId, clientId, sendUserId, code, msg, data)
			}
		}
	}
}

// 添加到本地分组
func (manager *ClientManager) AddClient2LocalGroup(groupName string, client *Client, userId string, extend string) {
	//标记当前客户端的userId
	client.UserId = userId
	client.Extend = extend

	//判断之前是否有添加过
	for _, groupValue := range client.GroupList {
		if groupValue == groupName {
			log.Printf("判断之前是否有添加过。。。。")
			return
		}
	}

	// 为属性添加分组信息
	groupKey := util.GenGroupKey(client.SystemId, groupName)
	//检查是超群组上限
	cCount := manager.GroupCount(groupKey)
	msg := Msg{}
	cId, _ := strconv.Atoi(groupName)
	set, _ := msg.GetSet(cId)
	log.Printf("配置%v,在线人数%d", set, cCount)
	if cCount >= set.OnLine {
		//当前用户下线，发送消息

		online := make(map[string]int)
		online["online"] = cCount
		online["is_show_online"] = set.IsShowOnline
		SendMessage2Client(client.ClientId, userId, retcode.OFFLINE_MESSAGE_CODE, "已达上线", online)
		time.Sleep(2 * time.Second)
		CloseClient(client.ClientId, client.SystemId)
		return
	}
	manager.addClient2Group(groupKey, client)

	client.GroupList = append(client.GroupList, groupName)
	log.Printf("群组列表client.GroupList%v\n", client.GroupList)
	//mJson, _ := json.Marshal(map[string]string{
	//	"clientId": client.ClientId,
	//	"userId":   client.UserId,
	//	"extend":   client.Extend,
	//})
	//data := string(mJson)
	//sendUserId := ""

	//发送系统通知
	if extend == "bullet_chat" {
		//SendMessage2Group(client.SystemId, sendUserId, groupName, retcode.ONLINE_MESSAGE_CODE, "", &data)
		retData := make(map[string]string)
		retData["chatroom_id"] = groupName
		retData["client_id"] = client.ClientId
		SendMessage2Client(client.ClientId, userId, retcode.ONLINE_MESSAGE_CODE, "客户端上线", retData)
	} else {
		//发送首页数据
		SendListMsgToClient(client.ClientId, userId, groupName)
	}
}

// 添加到本地分组
func (manager *ClientManager) addClient2Group(groupKey string, client *Client) {
	manager.GroupLock.Lock()
	defer manager.GroupLock.Unlock()
	manager.Groups[groupKey] = append(manager.Groups[groupKey], client.ClientId)
	//log.Printf("本地群组client.GroupList%v\n", manager.Groups)
}

// 删除分组里的客户端
func (manager *ClientManager) delGroupClient(groupKey string, clientId string) {
	manager.GroupLock.Lock()
	defer manager.GroupLock.Unlock()

	for index, groupClientId := range manager.Groups[groupKey] {
		if groupClientId == clientId {
			manager.Groups[groupKey] = append(manager.Groups[groupKey][:index], manager.Groups[groupKey][index+1:]...)
		}
	}
}

// 获取本地分组的成员
func (manager *ClientManager) GetGroupClientList(groupKey string) []string {
	manager.GroupLock.RLock()
	//log.Printf("所有分组用户%#v", manager.Groups[groupKey])
	cid := make([]string, 0, len(manager.Groups[groupKey]))
	for _, v := range manager.Groups[groupKey] {
		cid = append(cid, v)
	}
	manager.GroupLock.RUnlock()
	return cid
}

// 添加到系统客户端列表
func (manager *ClientManager) AddClient2SystemClient(systemId string, client *Client) {
	manager.SystemClientsLock.Lock()
	defer manager.SystemClientsLock.Unlock()
	manager.SystemClients[systemId] = append(manager.SystemClients[systemId], client.ClientId)
}

// 删除系统里的客户端
func (manager *ClientManager) delSystemClient(client *Client) {
	manager.SystemClientsLock.Lock()
	defer manager.SystemClientsLock.Unlock()

	for index, clientId := range manager.SystemClients[client.SystemId] {
		if clientId == client.ClientId {
			manager.SystemClients[client.SystemId] = append(manager.SystemClients[client.SystemId][:index], manager.SystemClients[client.SystemId][index+1:]...)
		}
	}
}

// 获取指定系统的客户端列表
func (manager *ClientManager) GetSystemClientList(systemId string) []string {
	manager.SystemClientsLock.RLock()
	defer manager.SystemClientsLock.RUnlock()
	return manager.SystemClients[systemId]
}
