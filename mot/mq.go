package mot

import (
	"github.com/motclub/common/mq"
	"time"
)

func XAddMessage(topic string, values map[string]interface{}) error {
	return Client.mq.XAddMessage(topic, values)
}

func XRegisterConsumer(topics []string, handler func(error, string, *mq.XMessage) error) {
	Client.mq.XRegisterConsumer(topics, handler)
}

func XRegisterConsumerGroup(topics []string, group, consumer string, handler func(error, string, *mq.XMessage) error) {
	Client.mq.XRegisterConsumerGroup(topics, group, consumer, handler)
}

func XAck(topic, group string, ids ...string) error {
	return Client.mq.XAck(topic, group, ids...)
}

func XDel(topic string, ids ...string) error {
	return Client.mq.XDel(topic, ids...)
}

func XRange(topic, start, end string, count int64) ([]mq.XMessage, error) {
	return Client.mq.XRange(topic, start, end, count)
}

func XPending(topic, group string, args *mq.XPendingArgs) ([]mq.XPendingReplyEntry, error) {
	return Client.mq.XPending(topic, group, args)
}

func XClaim(topic, group, consumer string, minIdle time.Duration, messages ...string) error {
	return Client.mq.XClaim(topic, group, consumer, minIdle, messages...)
}

func XExclusiveRegisterConsumer(topics []string, consumer string, handler func(error, string, *mq.XMessage) error) {
	Client.mq.XExclusiveRegisterConsumer(topics, consumer, handler)
}

func XExclusiveAck(topic string, ids ...string) error {
	return Client.mq.XExclusiveAck(topic, ids...)
}

func XExclusivePending(topic string, args *mq.XPendingArgs) ([]mq.XPendingReplyEntry, error) {
	return Client.mq.XExclusivePending(topic, args)
}

func XExclusiveClaim(topic, consumer string, minIdle time.Duration, messages ...string) error {
	return Client.mq.XExclusiveClaim(topic, consumer, minIdle, messages...)
}
