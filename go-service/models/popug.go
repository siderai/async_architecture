package models

import (
    "time"
)

type PopugModel struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"size:255;not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
}