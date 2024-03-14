package models

type Avatar struct {
	ID          int64  `gorm:"primaryKey" json:"id"`
	Image       string `gorm:"type:varchar(300)" json:"image"`
	Avatar_name string `gorm:"type:varchar(300)" json:"avatar_name"`
	Price       string `gorm:"type:varchar(300)" json:"price"`
}
