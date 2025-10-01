package models

import "gorm.io/gorm"

type Person struct {
	ID     uint64   `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	Name   string   `gorm:"index;column:name"`
	Cars   []*Car   `gorm:"foreignKey:OwnerID"`
	Movies []*Movie `gorm:"foreignKey:OwnerID"`
}

func (x Person) PreloadCars(db *gorm.DB) *gorm.DB {
	return db
}

func (x Person) PreloadMovies(db *gorm.DB) *gorm.DB {
	return db.Where("movie.isSeries = ?", false)
}
