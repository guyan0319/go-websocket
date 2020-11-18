package cache

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go-websocket/servers/msgs"
	"strconv"
)

const (
	serversHashKey       = "ws:hash:allservers" // 全部的服务器
	serversHashCacheTime = 2 * 60 * 60        // key过期时间
	serversHashTimeout   = 3 * 60             // 超时时间
)

func getServersHashKey() (key string) {
	key = fmt.Sprintf("%s", serversHashKey)

	return
}

// 设置服务器信息
func SetServerInfo(server *msgs.Server, currentTime uint64) (err error) {
	key := getServersHashKey()

	value := fmt.Sprintf("%d", currentTime)

	rc := RedisClient.Get()
	number, err := redis.Int(rc.Do("hSet", key, server.String(), value))
	if err != nil {
		fmt.Println("SetServerInfo", key, number, err)

		return
	}

	if number != 1 {

		return
	}

	rc.Do("Expire", key, serversHashCacheTime)

	return
}

// 下线服务器信息
func DelServerInfo(server *msgs.Server) (err error) {
	key := getServersHashKey()
	redisClient := RedisClient.Get()
	number, err := redisClient.Do("hDel", key, server.String())
	if err != nil {
		fmt.Println("DelServerInfo", key, number, err)

		return
	}

	if number != 1 {

		return
	}

	redisClient.Do("Expire", key, serversHashCacheTime)

	return
}

func GetServerAll(currentTime uint64) (servers []*msgs.Server, err error) {

	servers = make([]*msgs.Server, 0)
	key := getServersHashKey()

	// 从池里获取连接
	rc := RedisClient.Get()
	// 用完后将连接放回连接池
	defer rc.Close()

	serverMap, err := redis.StringMap(rc.Do("hGetAll", key))

	valByte, _ := json.Marshal(serverMap)
	fmt.Println("GetServerAll", key, string(valByte))

	for key, value := range serverMap {
		valueUint64, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			fmt.Println("GetServerAll", key, err)

			return nil, err
		}

		// 超时
		if valueUint64+serversHashTimeout <= currentTime {
			continue
		}

		server, err := msgs.StringToServer(key)
		if err != nil {
			fmt.Println("GetServerAll", key, err)

			return nil, err
		}

		servers = append(servers, server)
	}

	return
}
