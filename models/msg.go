package models

import (
	"fmt"
	. "github.com/woodylan/go-websocket/config"
	"log"
	"strings"
)

type Msg struct {
	Id          int         `json:"id"`
	ChatroomId  int         `json:"chatroom_id"`
	UserId      int         `json:"user_id"`
	Content     string      `json:"content"`
	ContentType int         `json:"content_type"`
	VideoCover  string      `json:"video_cover"`
	QuoteId     int         `json:"quote_id"`
	CreateTime  string      `json:"create_time"`
	User        UserType    `json:"user"`
	Quote       interface{} `json:"quote"`
	TimeMsg     interface{} `json:"time_msg"`
}
type BulletChat struct {
	ChatroomId int    `json:"chatroom_id"`
	Color      string `json:"color"`
	Content    string `json:"content"`
	ChatType   string `json:"chat_type"`
	ClientId   string `json:"client_id"`
}

type QuoteNull struct {
}

type QuoteType struct {
	Id          int      `json:"id"`
	UserId      int      `json:"user_id"`
	Content     string   `json:"content"`
	ContentType int      `json:"content_type"`
	VideoCover  string   `json:"video_cover"`
	User        UserType `json:"user"`
}
type UserType struct {
	UserId   int    `json:"user_id"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
}
type Subs struct {
	ChatroomId   int    `json:"chatroom_id"`
	LastId       int    `json:"last_id"`
	Page         int    `json:"page"`
	TopId        int    `json:"top_id"`
	OnLine       int    `json:"on_line"`
	PushMsgType  int    `json:"push_msg_type"`
	ClientId     string `json:"client_id"`
	IsShowOnline int    `json:"is_show_online"`
}
type Lists struct {
	List interface{} `json:"list"`
	Sub  Subs        `json:"sub"`
}
type MsgType struct {
	Msg Msg  `json:"msg_item"`
	Sub Subs `json:"sub"`
}
type ListNull struct {
}
type TimeMsg struct {
	Content string `json:"content"`
}
type TimeMsgNUll struct {
}

func (msg *Msg) GetList(lastId int) (msgs []Msg, err error) {
	sql := fmt.Sprintf("select id,chatroom_id,user_id,content,content_type,video_cover,quote_id,create_time from chatroom_msg where  is_delete=0 and report_count<2  order by id desc limit %d", ViperConfig.App.Limit)
	if lastId > 0 {
		sql = fmt.Sprintf("select id,chatroom_id,user_id,content,content_type,video_cover,quote_id,create_time from chatroom_msg where  id<%d and is_delete=0 and report_count<2  order by id desc limit  %d", lastId, ViperConfig.App.Limit)
	}
	//rows, err := Db.Query("select id,chatroom_id,user_id,content,content_type,quote_id,create_time from chatroom_msg where  is_delete=0 and report_count<2  order by id desc limit 30")
	log.Printf(sql)
	rows, err := Db.Query(sql)
	if err != nil {
		return
	}
	var ids_arr []int
	var qids_arr []int

	defer rows.Close()
	for rows.Next() {
		ms := Msg{}
		if err = rows.Scan(&ms.Id, &ms.ChatroomId, &ms.UserId, &ms.Content, &ms.ContentType, &ms.VideoCover, &ms.QuoteId, &ms.CreateTime); err != nil {
			return
		}
		ids_arr = append(ids_arr, ms.UserId)
		ms.Quote = QuoteNull{}
		if ms.QuoteId > 0 {
			qids_arr = append(qids_arr, ms.QuoteId)
		}
		if ms.ContentType > 0 {
			ms.Content = ViperConfig.App.OssUrl + ms.Content
		}
		if ms.ContentType == 3 && ms.VideoCover != "" {
			ms.VideoCover = ViperConfig.App.OssUrl + ms.VideoCover
		}

		msgs = append(msgs, ms)

	}
	if len(msgs) == 0 {
		return
	}

	ids := strings.Replace(strings.Trim(fmt.Sprint(ids_arr), "[]"), " ", ",", -1)
	qids := strings.Replace(strings.Trim(fmt.Sprint(qids_arr), "[]"), " ", ",", -1)
	uidsMap, _ := msg.GetUserByIds(ids)
	qidsMap := make(map[int]QuoteType)
	if qids != "" {
		qidsMap, _ = msg.GetQuoteByIds(qids)
	}
	for k, v := range msgs {
		msgs[k].User = uidsMap[v.UserId]
		if v.QuoteId > 0 {
			msgs[k].Quote = qidsMap[v.QuoteId]
		} else {
			msgs[k].Quote = QuoteNull{}
		}
	}
	return
}

func (msg *Msg) GetMsg(id int) (ms Msg, err error) {
	//log.Printf("读取单条消息")
	stout, err := Db.Prepare("select id,chatroom_id,user_id,content,content_type,quote_id,create_time from chatroom_msg where id=? and is_delete=0 ")
	if err != nil {
		return
	}
	defer stout.Close()
	err = stout.QueryRow(id).Scan(&ms.Id, &ms.ChatroomId, &ms.UserId, &ms.Content, &ms.ContentType, &ms.QuoteId, &ms.CreateTime)
	return
}
func (msg *Msg) GetTop() (id int, err error) {
	//log.Printf("读取置顶数据消息")
	stout, err := Db.Prepare("select id from chatroom_msg where is_top=1 and is_delete=0 and report_count<2  order by update_time desc  limit 1 ")
	if err != nil {
		log.Printf("置顶查询出错了吗？%v\n", err)
		return
	}
	defer stout.Close()
	err = stout.QueryRow().Scan(&id)
	if err != nil {
		log.Printf("置顶Scan查询出错了吗？%v\n", err)
		return
	}
	return
}

func (msg *Msg) GetUser(uid int) (u UserType, err error) {
	//log.Printf("读取数uid")
	stout, err := Db.Prepare("select id,nick_name,avatar from user where id=? ")
	if err != nil {
		return
	}
	defer stout.Close()
	err = stout.QueryRow(uid).Scan(&u.UserId, &u.NickName, &u.Avatar)
	return
}
func (msg *Msg) GetUserByIds(ids string) (myMap map[int]UserType, err error) {
	//log.Printf("读取数uid")
	sql := fmt.Sprintf("select id,nick_name,avatar from user where id in (%s)", ids)
	log.Printf(sql)
	rows, err := Db.Query(sql)
	defer rows.Close()
	if err != nil {
		return
	}
	var u []UserType
	for rows.Next() {
		ut := UserType{}
		if err = rows.Scan(&ut.UserId, &ut.NickName, &ut.Avatar); err != nil {
			return
		}
		ut.Avatar = ViperConfig.App.OssUrl + ut.Avatar
		u = append(u, ut)
	}
	//log.Printf("用户数据%v\n", u)
	myMap = make(map[int]UserType)
	for _, v := range u {
		myMap[v.UserId] = v
	}
	return
}
func (msg *Msg) GetQuoteByIds(ids string) (myMap map[int]QuoteType, err error) {
	//log.Printf("读取数uid")
	sql := fmt.Sprintf("select id,user_id,content,content_type,video_cover from chatroom_msg where id in (%s) and is_delete=0", ids)
	log.Printf(sql)
	rows, err := Db.Query(sql)
	defer rows.Close()
	if err != nil {
		return
	}
	var quids_arr []int
	var u []QuoteType
	for rows.Next() {
		ms := QuoteType{}
		if err = rows.Scan(&ms.Id, &ms.UserId, &ms.Content, &ms.ContentType, &ms.VideoCover); err != nil {
			return
		}
		if ms.ContentType > 0 {
			ms.Content = ViperConfig.App.OssUrl + ms.Content
		}
		quids_arr = append(quids_arr, ms.UserId)
		u = append(u, ms)
	}

	uids := strings.Replace(strings.Trim(fmt.Sprint(quids_arr), "[]"), " ", ",", -1)
	uidsMap, _ := msg.GetUserByIds(uids)
	myMap = make(map[int]QuoteType)
	//log.Printf("引用内的uids %#v\n", uidsMap)
	for k, v := range u {
		u[k].User = uidsMap[v.UserId]
	}
	for _, v := range u {
		myMap[v.Id] = v
	}
	//log.Printf("引用内的数据 %#v\n", myMap)
	return
}
