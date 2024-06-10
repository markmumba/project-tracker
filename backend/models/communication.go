package models

import (
    "gorm.io/gorm"
)

type CommunicationHistory struct {
    gorm.Model
    ProjectID   uint   `json:"project_id"`
    UserID      uint   `json:"user_id"`
    Message     string `json:"message"`
    MessageDate string `json:"message_date"`
    Project     Project `gorm:"foreignKey:ProjectID"`
    User        User    `gorm:"foreignKey:UserID"`
}
