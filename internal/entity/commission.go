package entity

type Commission struct {
	ID     uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Type   string  `gorm:"not null;varchar(255)" json:"type"`
	Title  string  `gorm:"not null;varchar(255)" json:"title"`
	Rate   float32 `gorm:"not null;type:decimal(15,2)" json:"rate"`
	Status bool    `gorm:"not null" json:"status"`
}