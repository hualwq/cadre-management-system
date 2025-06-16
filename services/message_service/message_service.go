package message_service

import (
	"context"
	"fmt"

	"cadre-management/pkg/setting"

	"github.com/segmentio/kafka-go"
)

// SendMessage 发送消息到 Kafka
func SendMessage(recipientID string, message string) error {
	// 创建 Kafka 写入器
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: setting.AppSetting.Kafka.Brokers,
		Topic:   setting.AppSetting.Kafka.Topic,
	})

	// 构建消息
	msg := kafka.Message{
		Key:   []byte(recipientID),
		Value: []byte(message),
	}

	// 发送消息
	err := w.WriteMessages(context.Background(), msg)
	if err != nil {
		return fmt.Errorf("failed to send message to Kafka: %w", err)
	}

	// 关闭写入器
	return w.Close()
}

// SubscribeMessages 订阅 Kafka 消息
func SubscribeMessages() {
	// 创建 Kafka 读取器
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: setting.AppSetting.Kafka.Brokers,
		Topic:   setting.AppSetting.Kafka.Topic,
		GroupID: "cadre-group",
	})

	defer r.Close()

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("failed to read message from Kafka: %v\n", err)
			continue
		}

		recipientID := string(msg.Key)
		message := string(msg.Value)
		fmt.Printf("Received message for %s: %s\n", recipientID, message)
	}
}

