package models

type Bookmark struct {
	InteractorModel
	Source User `gorm:"foreignKey:UserID"`
	Target Post `gorm:"foreignKey:PostID"`
}
