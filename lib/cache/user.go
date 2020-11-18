

package cache

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"go-websocket/servers/msgs"
)

const (
	userOnlinePrefix    = "user:online:" // 用户在线状态
	userOnlineCacheTime = 24 * 60 * 60
)

/*********************  查询用户是否在线  ************************/
func getUserOnlineKey(userKey string) (key string) {
	key = fmt.Sprintf("%s%s", userOnlinePrefix, userKey)

	return
}

func GetUserOnlineInfo(userKey string) (userOnline *msgs.UserOnline, err error) {

	key := getUserOnlineKey(userKey)

	data, err := GetBytes(key)
	if err != nil {
		if err == redis.Nil {
			fmt.Println("GetUserOnlineInfo", userKey, err)

			return
		}

		fmt.Println("GetUserOnlineInfo", userKey, err)

		return
	}

	userOnline = &msgs.UserOnline{}
	err = json.Unmarshal(data, userOnline)
	if err != nil {
		fmt.Println("获取用户在线数据 json Unmarshal", userKey, err)

		return
	}

	fmt.Println("获取用户在线数据", userKey, "time", userOnline.LoginTime, userOnline.HeartbeatTime, "AccIp", userOnline.AccIp, userOnline.IsLogoff)

	return
}

// 设置用户在线数据
func SetUserOnlineInfo(userKey string, userOnline *msgs.UserOnline) (err error) {

	redisClient := RedisClient.Get()

	key := getUserOnlineKey(userKey)

	valueByte, err := json.Marshal(userOnline)
	if err != nil {
		fmt.Println("设置用户在线数据 json Marshal", key, err)

		return
	}

	_, err = redisClient.Do("setEx", key, userOnlineCacheTime, string(valueByte))
	if err != nil {
		fmt.Println("设置用户在线数据 ", key, err)

		return
	}

	return
}
