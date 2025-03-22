package rocketmqKit

import (
	"context"
	"fmt"
	rmqClient "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
	"time"
)

// InitRocketMQConsumers 初始化消费者
/*
 * @param ctx {context.Context} 上下文
 * @param maxMessageNum {int32} 最大消息数
 * @param invisibleDuration {int} 持续时间·单位s
 */
func (r *RocketMQConfig) InitRocketMQConsumers() error {
	var err error
	RocketMQConsumer, err = rmqClient.NewSimpleConsumer(&rmqClient.Config{
		Endpoint:      r.Endpoint,
		NameSpace:     r.Namespace,
		ConsumerGroup: r.GroupName,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    r.AccessKey,
			AccessSecret: r.SecretKey,
		},
	},
		rmqClient.WithAwaitDuration(time.Second*time.Duration(r.WaitTime)),
	)
	if err != nil {
		return err
	}
	err = RocketMQConsumer.Start()
	if err != nil {
		return err
	}
	return nil
}

// Subscription 消息订阅 *自定义程度过高，目前返回通用的Any数据包，自行解包
func Subscription(ctx context.Context, maxMessageNum int32, invisibleDuration int) error {
	// 订阅 Topic
	err := RocketMQConsumer.Subscribe(publicTopic, rmqClient.NewFilterExpression("*"))
	if err != nil {
		return err
	}
	for {
		mvs, _ := RocketMQConsumer.Receive(ctx, maxMessageNum, time.Duration(invisibleDuration)*time.Second)
		// ack message
		for _, mv := range mvs {
			// 打印包，暂时没想好怎么回传
			fmt.Println(mv)
			if true {
				// 如果消费成功进行应答
				RocketMQConsumer.Ack(ctx, mv)
			}
		}
	}
}
