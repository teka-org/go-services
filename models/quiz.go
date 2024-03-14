package models

type Quiz struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	Question string `gorm:"type:varchar(300)" json:"question"`
	Answer   string `gorm:"type:varchar(300)" json:"answer"`
}
