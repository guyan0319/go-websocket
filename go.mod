module go-websocket

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/gomodule/redigo v1.8.3
	github.com/gorilla/websocket v1.4.2
	github.com/pkg/errors v0.9.1
	github.com/rpcxio/rpcx-examples v1.1.6
	github.com/smallnest/rpcx v0.0.0-20201201225933-a922ef376871
	google.golang.org/grpc/examples v0.0.0-20201130222003-4a0125ac5808 // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.8
)
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0