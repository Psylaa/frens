package models

type Follow struct {
	InteractorModel
	Source User `gorm:"foreignKey:UserID"`
	Target User `gorm:"foreignKey:UserID"`
}
