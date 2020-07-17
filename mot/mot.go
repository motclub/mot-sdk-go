package mot

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/motclub/common/caller"
	"github.com/motclub/common/std"
	"github.com/robfig/cron/v3"
	"os"
	"strings"
	"time"

	_ "github.com/motclub/common/caller/http"
	_ "github.com/motclub/common/caller/rpcx"

	"github.com/motclub/common/cache"
	redisCache "github.com/motclub/common/cache/redis"

	"github.com/motclub/common/mq"
	redisMQ "github.com/motclub/common/mq/redis"
)

var Client = &client{
	appId:            "default",
	instanceId:       uuid.New().String(),
	cron:             cron.New(cron.WithSeconds()),
	HeartbeatPeriod:  10 * time.Second,
	HeartbeatTimeout: 30 * time.Second,
}

type client struct {
	appId             string
	instanceId        string
	coreServerAddress string
	moduleCode        string
	moduleConfig      std.D
	cache             cache.ICache
	mq                mq.IMessageQueue
	cron              *cron.Cron

	HeartbeatPeriod  time.Duration
	HeartbeatTimeout time.Duration
}

func (c *client) instanceKey() string {
	return fmt.Sprintf("mot/%s/%s/%s", c.appId, c.moduleCode, c.instanceId)
}

func (c *client) runJobs() error {
	spec := fmt.Sprintf("@every %s", c.HeartbeatPeriod.String())
	if _, err := c.cron.AddFunc(spec, func() {
		_ = c.cache.Set(c.instanceKey(), time.Now().Unix(), 3*c.HeartbeatPeriod)
	}); err != nil {
		return err
	}
	return nil
}

func New(moduleCode string, coreServerAddress ...string) error {
	// 核心模块地址
	var server string
	if v := os.Getenv("MOT_SERVER"); v != "" {
		server = v
	} else if len(coreServerAddress) > 0 && coreServerAddress[0] != "" {
		server = coreServerAddress[0]
	} else {
		server = "tcp@mot:8889"
	}
	if v := os.Getenv("MOT_MODULE_CODE"); v != "" {
		moduleCode = v
	}
	Client.moduleCode = moduleCode
	if v := os.Getenv("MOT_APP_ID"); v != "" {
		Client.appId = v
	}
	if !strings.Contains(server, "@") && !strings.HasPrefix(server, "http") {
		server = fmt.Sprintf("tcp@%s", server)
	}
	Client.coreServerAddress = server
	// 查询模块配置
	reply := caller.Call(&caller.ServiceEntry{
		Kind: "RPCX",
		Settings: std.D{
			"server":         Client.coreServerAddress,
			"service_path":   "core",
			"service_method": "config",
		},
	}, &std.Args{
		RequestID:      uuid.New().String(),
		SessionPayload: nil,
		Data:           std.D{"mod": moduleCode},
	})
	if err := reply.Bind(&Client.moduleConfig); err != nil {
		panic(err)
	}
	config := Client.moduleConfig
	// 初始化缓存
	switch adapter := config.GetString("cache_adapter"); adapter {
	case "redis":
		var opts redis.UniversalOptions
		if !config.HasGet("cache_configs", &opts) {
			panic(errors.New("failed to get cache configuration"))
		}
		if c, err := redisCache.NewRedisCache(&opts); err != nil {
			panic(err)
		} else {
			Client.cache = c
		}
	default:
		return fmt.Errorf("unsupported cache adapter: %s", adapter)
	}

	// 初始化消息队列
	switch adapter := config.GetString("mq_adapter"); adapter {
	case "redis":
		var opts redis.UniversalOptions
		if !config.HasGet("mq_configs", &opts) {
			panic(errors.New("failed to get message queue configuration"))
		}
		if c, err := redisMQ.NewRedisMQ(&opts); err != nil {
			panic(err)
		} else {
			Client.mq = c
		}
	default:
		return fmt.Errorf("unsupported message queue adapter: %s", adapter)
	}

	if err := Client.runJobs(); err != nil {
		panic(err)
	}
	return nil
}

func Close() error {
	var err error
	if Client.cache != nil {
		if e := Client.cache.Close(); e != nil {
			err = e
		}
	}
	if Client.mq != nil {
		if e := Client.mq.XClose(); e != nil {
			err = e
		}
	}
	return err
}
