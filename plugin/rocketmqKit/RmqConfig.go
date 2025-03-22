package rocketmqKit

import (
	rmqClient "github.com/apache/rocketmq-clients/golang/v5"
)

type RocketMQConfig struct {
	Topic     string `json:"topic"`     // Topic 主题
	Endpoint  string `json:"endpoint"`  // Endpoint 端点
	Namespace string `json:"nameSpace"` // 命名空间
	GroupName string `json:"groupName"` // 组名
	AccessKey string `json:"accessKey"` // AccessKey 访问密钥
	SecretKey string `json:"secretKey"` // SecretKey 密钥
	WaitTime  int    `json:"waitTime"`  // 等待时间
}

var (
	RocketMQProducer rmqClient.Producer       // 对外暴露RocketMq的生产者
	RocketMQConsumer rmqClient.SimpleConsumer // 对外暴露RocketMq的消费者
	publicTopic      string                   // 主题内部暴露
)
