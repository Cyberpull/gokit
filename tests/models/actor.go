package models

import (
	"gorm.io/gorm"
)

type Actor struct {
	ID   uint64 `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	Name string `gorm:"index;column:name"`
	Age  int    `gorm:"index;column:age"`
}

func (x *Actor) ScopeAdult(db *gorm.DB) *gorm.DB {
	return db.Where("age > ?", 18)
}
