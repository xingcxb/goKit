// rocketmq 生产者

package rocketmqKit

import (
	"context"
	"errors"
	"fmt"
	rmqClient "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
	"os"
)

// InitRocketMqProducers 初始化rocketMq生产者
func (r *RocketMQConfig) InitRocketMqProducers() error {
	// 设置日志级别
	_ = os.Setenv(rmqClient.ENABLE_CONSOLE_APPENDER, "false") // 关闭控制台日志
	_ = os.Setenv(rmqClient.CLIENT_LOG_LEVEL, "error")
	var err error
	if r.Endpoint == "" {
		return errors.New("rmq的Endpoint参数为空，请联系开发人员")
	}
	RocketMQProducer, err = rmqClient.NewProducer(
		&rmqClient.Config{
			Endpoint:      r.Endpoint,
			NameSpace:     r.Namespace,
			ConsumerGroup: r.GroupName,
			Credentials: &credentials.SessionCredentials{
				AccessKey:    r.AccessKey,
				AccessSecret: r.SecretKey,
			},
		},
		rmqClient.WithTopics(r.Topic),
	)
	if err != nil {
		return err
	}
	if RocketMQProducer == nil {
		return errors.New("创建rmq客户端失败，RmqProducer为空")
	}
	err = RocketMQProducer.Start()
	if err != nil {
		return err
	}

	return nil
}

// Send 发送消息
/*
 * @param ctx {context.Context} 上下文
 * @param msg {[]byte} 消息主体
 * @param key {string}
 * @param tag {string}
 * @param messageGroup {string}
 */
func Send(ctx context.Context, msg []byte, key, tag, messageGroup string) error {
	msgObj := &rmqClient.Message{
		Topic: publicTopic,
		Body:  msg,
	}
	msgObj.SetKeys(key)
	msgObj.SetTag(tag)
	msgObj.SetMessageGroup(messageGroup)
	_, err := RocketMQProducer.Send(ctx, msgObj)
	if err != nil {
		return errors.New(fmt.Sprintf("发送流量日志失败，程序关闭%s", err.Error()))
	}
	return nil
}
