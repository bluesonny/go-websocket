package util

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/woodylan/go-websocket/pkg/setting"
	"github.com/woodylan/go-websocket/tools/crypto"
	"strconv"
	"strings"
	"time"
)

//GenUUID 生成uuid
func GenUUID() string {
	uuidFunc := uuid.NewV4()
	uuidStr := uuidFunc.String()
	uuidStr = strings.Replace(uuidStr, "-", "", -1)
	uuidByt := []rune(uuidStr)
	return string(uuidByt[8:24])
}

//对称加密IP和端口，当做clientId
func GenClientId() string {
	raw := []byte(setting.GlobalSetting.LocalHost + ":" + setting.CommonSetting.RPCPort)
	str, err := crypto.Encrypt(raw, []byte(setting.CommonSetting.CryptoKey))
	if err != nil {
		panic(err)
	}

	return str
}

//解析redis的地址格式
func ParseRedisAddrValue(redisValue string) (host string, port string, err error) {
	if redisValue == "" {
		err = errors.New("解析地址错误")
		return
	}
	addr := strings.Split(redisValue, ":")
	if len(addr) != 2 {
		err = errors.New("解析地址错误")
		return
	}
	host, port = addr[0], addr[1]

	return
}

//判断地址是否为本机
func IsAddrLocal(host string, port string) bool {
	return host == setting.GlobalSetting.LocalHost && port == setting.CommonSetting.RPCPort
}

//是否集群
func IsCluster() bool {
	return setting.CommonSetting.Cluster
}

//获取client key地址信息
func GetAddrInfoAndIsLocal(clientId string) (addr string, host string, port string, isLocal bool, err error) {
	//解密ClientId
	addr, err = crypto.Decrypt(clientId, []byte(setting.CommonSetting.CryptoKey))
	if err != nil {
		return
	}

	host, port, err = ParseRedisAddrValue(addr)
	if err != nil {
		return
	}

	isLocal = IsAddrLocal(host, port)
	return
}

func GenGroupKey(systemId, groupName string) string {
	return systemId + ":" + groupName
}

/**
 * 微信展示时间的方法
 * @param $addTime
 * @return string
 */
func GetChatTimeStr(addTimeInt int64) (timeStr string) {
	nowTimeInt := time.Now().Unix()

	if addTimeInt > nowTimeInt {
		return
	}
	weekNum := map[string]int{"Sun": 0, "Mon": 1, "Tue": 2, "Wed": 3, "Thu": 4, "Fri": 5, "Sat": 6}
	am := map[string]int{"AM": 0, "PM": 1}
	//返回的时间
	timeStr = ""
	//获取当前时间
	t := time.Unix(int64(addTimeInt), 0)
	addTimeStr := t.Format("2006,1,2,Mon,PM,15,04,06")
	//log.Printf("addTimeStr------ %v", addTimeStr)
	addTimeArr := strings.Split(addTimeStr, ",")
	//log.Printf("addTimeArr %v", addTimeArr)
	var addTime []int
	for k, v := range addTimeArr {
		if k == 3 {
			w := weekNum[v]
			addTime = append(addTime, w)
		} else if k == 4 {
			a := am[v]
			addTime = append(addTime, a)
		} else {
			v1, _ := strconv.Atoi(v)
			addTime = append(addTime, v1)
		}
	}
	//log.Printf("addTime %v", addTime)

	t2 := time.Unix(int64(nowTimeInt), 0)
	nowTimeStr := t2.Format("2006,1,2,Mon,PM,15,04,06")
	nowTimeArr := strings.Split(nowTimeStr, ",")
	var nowTime []int
	for k, v := range nowTimeArr {
		if k == 3 {
			w := weekNum[v]
			nowTime = append(nowTime, w)
		} else if k == 4 {
			a := am[v]
			nowTime = append(nowTime, a)
		} else {
			v1, _ := strconv.Atoi(v)
			nowTime = append(nowTime, v1)
		}
	}
	//log.Printf("nowTime %v", nowTime)
	//$addTime = explode(',', date('Y,n,j,w,a,h,i,y', $addTime));//年，月，日，星期，上下午，时，分
	//$nowTime = explode(',', date('Y,n,j,w,a,h,i,y', $nowTime));

	dayPerMonthAddTime := getDayPerMonth(addTime[0])
	week := map[int]string{0: "星期日", 1: "星期一", 2: "星期二", 3: "星期三", 4: "星期四", 5: "星期五", 6: "星期六"}
	//如果时间差小于一天的,显示（上午 时间） / （下午 时间）
	if addTime[0] == nowTime[0] && addTime[1] == nowTime[1] && addTime[2] == nowTime[2] {
		timeStr = timeStr + addTimeArr[5] + ":" + addTimeArr[6]
	} else if (addTime[0] == nowTime[0] && addTime[1] == nowTime[1] && addTime[2] == nowTime[2]-1) || (addTime[0] == nowTime[0] && nowTime[1]-addTime[1] == 1 && dayPerMonthAddTime[addTime[1]] == addTime[2] && nowTime[2] == 1) || (nowTime[0]-addTime[0] == 1 && addTime[1] == 12 && addTime[2] == 31 && nowTime[1] == 1 && nowTime[2] == 1) { //如果时间差在昨天,三种情况（同一月份内跨一天、月末跨越到月初、年末跨越到年初）显示格式：昨天 时:分 上午/下午
		timeStr += "昨天 " + addTimeArr[5] + ":" + addTimeArr[6] + " "
	} else if (addTime[0] == nowTime[0] && addTime[1] == nowTime[1] && nowTime[2]-addTime[2] < 7) || (addTime[0] == nowTime[0] && nowTime[1]-addTime[1] == 1 && dayPerMonthAddTime[addTime[1]]-addTime[2]+nowTime[2] < 7 || (nowTime[0]-addTime[0] == 1 && addTime[1] == 12 && nowTime[1] == 1 && 31-addTime[2]+nowTime[2] < 7)) { //如果时间差在一个星期之内的,也是三种情况，显示格式：星期 时:分 上午/下午

		timeStr += week[addTime[3]] + " " + addTimeArr[5] + ":" + addTimeArr[6]
	} else { //显示格式：月/日/年 时:分 上午/下午
		timeStr += addTimeArr[1] + "/" + addTimeArr[2] + "/" + addTimeArr[7] + " " + addTimeArr[5] + ":" + addTimeArr[6]
	}

	if addTime[4] == 0 {
		//timeStr += " 上午"
	} else if addTime[4] == 1 {
		//timeStr += " 下午"
	}

	return

}

//根据年份获取每个月份的总天数和每年最后一个月的天数
func getDayPerMonth(year int) (m2 map[int]int) {
	m2 = make(map[int]int)

	m2[1] = 31
	m2[3] = 31
	m2[4] = 30
	m2[5] = 31
	m2[6] = 30
	m2[7] = 31
	m2[8] = 31
	m2[9] = 30
	m2[10] = 31
	m2[11] = 30
	m2[12] = 31

	//闰年
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		m2[2] = 29
	} else {
		m2[2] = 28
	}
	return
}
