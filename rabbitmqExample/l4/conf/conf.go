package config

const (
	RMQADDR     = "amqp://admin:123456@127.0.0.1:5672/"
	EXCHANGENAME = "syslog_direct"
	CONSUMERCNT  = 4
)

var (
	RoutingKeys [4]string = [4]string{"info", "debug", "warn", "error"}
)
