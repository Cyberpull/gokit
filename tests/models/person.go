package models

import "gorm.io/gorm"

type Person struct {
	ID     uint64   `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	Name   string   `gorm:"index;column:name"`
	Cars   []*Car   `gorm:"foreignKey:OwnerID" gokit-dbo:"preload"`
	Movies []*Movie `gorm:"foreignKey:OwnerID" gokit-dbo:"preload"`
}

func (x Person) MoviesPreloader(db *gorm.DB) *gorm.DB {
	return db.Where("movie.isSeries = ?", false)
}
