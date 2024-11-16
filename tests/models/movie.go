package models

type Movie struct {
	ID       uint64  `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	Name     string  `gorm:"index;column:name"`
	IsSeries bool    `gorm:"index;column:isSeries;not null;default:false"`
	OwnerID  uint64  `gorm:"index;column:ownerId"`
	Owner    *Person `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
