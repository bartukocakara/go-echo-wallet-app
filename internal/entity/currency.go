package entity

type Currency struct {
	ID   uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Code string `gorm:"not null;varchar(255)" json:"code"`
}