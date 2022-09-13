module github.com/lucky-cheerful-man/phoenix_gateway

go 1.16

require (
	github.com/astaxie/beego v1.12.3
	github.com/gin-gonic/gin v1.8.1
	github.com/go-ini/ini v1.67.0
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/lucky-cheerful-man/phoenix_apis v1.0.0
	github.com/matoous/go-nanoid v1.5.0
	github.com/micro/micro v1.18.0 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/nats-io/nats-server/v2 v2.9.0 // indirect
	github.com/sirupsen/logrus v1.9.0
	google.golang.org/grpc v1.49.0 // indirect
	google.golang.org/grpc/examples v0.0.0-20220908195427-552de12024bc // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace github.com/micro/micro v1.18.0 => github.com/micro/micro v1.10.0
