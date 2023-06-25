package database

import "github.com/google/uuid"

// Status represents a status update by a user.
type Status struct {
	BaseModel
	UserID uuid.UUID `json:"userId"`
	Text   string    `json:"text"`
	Media  []Media   `gorm:"ForeignKey:StatusID"`
}

// GetStatus gets a status update by ID.
func GetStatus(id uuid.UUID) (*Status, error) {
	var status Status
	if err := db.Preload("Media").First(&status, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &status, nil
}

// CreateStatus creates a new status update.
func CreateStatus(userID uuid.UUID, text string, media []Media) (*Status, error) {
	newStatus := Status{
		BaseModel: BaseModel{ID: uuid.New()},
		UserID:    userID,
		Text:      text,
		Media:     media,
	}

	if err := db.Create(&newStatus).Error; err != nil {
		return nil, err
	}

	return &newStatus, nil
}

func DeleteStatus(statusID uuid.UUID) error {
	err := db.Delete(&Status{}, "id = ?", statusID).Error
	return err
}
