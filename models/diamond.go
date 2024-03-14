package models

type Diamond struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	Quantity int    `gorm:"type:varchar(300)" json:"quantity"`
	Image    string `gorm:"type:varchar(300)" json:"image"`
	Price    int    `gorm:"type:varchar(300)" json:"price"`
}
