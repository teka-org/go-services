package models

type User struct {
	ID    int64  `gorm:"primaryKey" json:"id"`
	Image string `gorm:"type:varchar(300)" json:"image"`
	Name  string `gorm:"type:varchar(300)" json:"name"`
	Email string `gorm:"type:varchar(300)" json:"email"`
}
