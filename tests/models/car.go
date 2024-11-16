package models

type Car struct {
	ID      uint64  `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	Brand   string  `gorm:"index;column:brand"`
	Color   string  `gorm:"index;column:color"`
	OwnerID uint64  `gorm:"index;column:ownerId"`
	Owner   *Person `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
