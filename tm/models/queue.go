package models

import (
	"time"
)

type Queue struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Category  string    //1数据库 2file
	Status    int       //1未执行 2执行中 3为执行完成 4为执行失败
	Content   string    //内容
	Report    string    //反馈结果
	Fn        string    //对应方法
	Message   string    //备注
}

func Report(id string, report string, message string, status int) {
	var queue Queue
	DB.Model(&queue).Where("id = ?", id).Updates(Queue{Report: report, Message: message, Status: status})
}

func Active(id string, status string) {
	var queue Queue
	DB.Model(&queue).Where("id = ?", id).Update("status", status)
}

func Size() int64 {
	var count int64
	DB.Model(&Queue{}).Where("status = ?", 1).Count(&count)
	return count
}

func Get(id int) Queue {
	var queue Queue
	DB.Where("id = ?", id).Order("created_at desc").First(&queue)
	return queue
}

func Pop() []Queue {
	var queues []Queue
	DB.Where("status = ?", 1).Order("created_at desc").Find(&queues)
	for _, queue := range queues {
		Active(string(queue.ID), "2")
	}
	return queues
}

func Push(category string, content string) {
	queue := &Queue{Category: category, Content: content, Status: 1}
	DB.Create(queue)
}

func init() {
	DB.AutoMigrate(&Queue{})
}
