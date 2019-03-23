package models

import (
	"time"
)

type Message struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserID    uint      `json:"userId"`
	Name      string    `json:"name" gorm:"-"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

func (message Message) TableName() string {
	return "message"
}

func (message *Message) Save(userId uint, userName string) {
	if len(message.Message) < 1 {
		return
	}

	message.UserID = userId
	message.Name = userName
	message.CreatedAt = time.Now().UTC()

	GetDB().Create(message)
}

func GetAllMessage() []*Message {
	messages := []*Message{}
	GetDB().Table("message").Select("message.*, user.name").Joins("INNER JOIN user ON user.id = message.user_id").Scan(&messages)

	return messages
}
