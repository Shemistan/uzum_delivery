package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// Настройка подключения к кафке
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",   // адресс
		"group.id":          "my-consumer-group", // некая группа консьюмеров
		"auto.offset.reset": "earliest",          // с какого момента читать сообщения
	}

	// Создаем консьюмера
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		log.Println("failed to create consumer")
		return
	}

	topics := []string{"my_topic"}

	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		log.Println("failed to send tipics")
		return
	}

	log.Println("Жду сообщений----------->")
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			// Анмаршалим то что пришло, кидаем это в структуру Order и после этого кидаем в местод Add()  который
			// сохранит это в БД
			fmt.Println(string(msg.Value)) // service.Add()
		} else {
			fmt.Println("Error to read message")
		}
	}
}
