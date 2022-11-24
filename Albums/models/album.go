package models

type album struct {
	ID     string  `gorm:"primary key" json:"id"`
	Title  string  `gorm:"type:varchar(100)" json:"title"`
	Artist string  `gorm:"type:varchar(300)" json:"artist"`
	Price  float64 `json:"price"`
}
