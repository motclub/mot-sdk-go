module github.com/motclub/mot-sdk-go

go 1.14

require (
	github.com/go-redis/redis/v8 v8.0.0-beta.6
	github.com/google/uuid v1.1.1
	github.com/motclub/common v0.0.0-20200717111949-59920611d82e
	github.com/robfig/cron/v3 v3.0.1
	github.com/sirupsen/logrus v1.6.0
	github.com/smallnest/rpcx v0.0.0-20200714114247-35b07de0def7 // indirect
)

// replace  github.com/motclub/common => ../common
