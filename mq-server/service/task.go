package service

import (
	"encoding/json"
	"fmt"
	"mq-server/model"
)

func CreateTask() {
	channel, err := model.MQ.Channel()
	if err != nil {
		panic(err)
	}

	declare, _ := channel.QueueDeclare("task_queue", true, false, false, false, nil)
	err = channel.Qos(1, 0, false)
	consume, err := channel.Consume(declare.Name, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	//一直监听队列
	go func() {
		for delivery := range consume {
			var t model.Task
			_ = json.Unmarshal(delivery.Body, &t)
			model.DB.Create(&t)
			fmt.Println("Done")
			_ = delivery.Ack(false)
		}
	}()

}
