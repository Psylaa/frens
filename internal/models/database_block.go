package models

type Block struct {
	InteractorModel
	Source User `gorm:"foreignKey:UserID"`
	Target User `gorm:"foreignKey:UserID"`
}
