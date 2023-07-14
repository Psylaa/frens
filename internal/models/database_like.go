package models

type Like struct {
	InteractorModel
	Source User `gorm:"foreignKey:UserID"`
	Target Post `gorm:"foreignKey:PostID"`
}
