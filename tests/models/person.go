package models

type Person struct {
	ID   uint64 `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	Name string `gorm:"index;column:name"`
	Cars []*Car `gorm:"foreignKey:OwnerID" gokit-dbo:"preload"`
}
