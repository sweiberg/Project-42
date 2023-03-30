package model

import (
	"exert-shop/db"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model

	ParentID   uint      `gorm:"not null" json:"parentID"`
	ReceiverID uint      `gorm:"not null" json:"receiverID"`
	SenderID   uint      `gorm:"not null" json:"senderID"`
	Message    string    `gorm:"type:text;not null" json:"message"`
	Subject    string    `gorm:"not null" json:"subject"`
	Replies    []Message `gorm:"foreignKey:ParentID"`
	Sender     User      `gorm:"foreignKey:SenderID"`
	Receiver   User      `gorm:"foreignKey:ReceiverID"`
}

func (message *Message) Create() (*Message, error) {
	err := db.Database.Create(&message).Error

	if err != nil {
		return &Message{}, err
	}

	return message, nil
}

func GetMessageByID(id uint64) (Message, error) {
	var message Message

	err := db.Database.Preload("Sender").Preload("Receiver").Where("id=?", id).Find(&message).Error

	if err != nil {
		return Message{}, err
	}

	return message, nil
}
