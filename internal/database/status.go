package database

import (
	"time"

	"github.com/google/uuid"
)

// Status represents a status update by a user.
type Status struct {
	BaseModel
	UserID uuid.UUID `json:"userId" gorm:"column:user_id;not null"`
	User   User      `json:"user" gorm:"foreignKey:UserID"`
	Text   string    `json:"text"`
	Media  []Media   `json:"media" gorm:"ForeignKey:StatusID"`
}

// GetStatus gets a status update by ID.
func GetStatus(id uuid.UUID) (*Status, error) {
	var status Status
	if err := db.Preload("User").Preload("Media").First(&status, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &status, nil
}

// GetStatusesByUserID gets all status updates by a user.
func GetStatusesByUserID(userID uuid.UUID) ([]Status, error) {
	var statuses []Status
	if err := db.Preload("User").Preload("Media").
		Order("created_at desc").
		Find(&statuses, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return statuses, nil
}

// GetStatusesByUserIDs gets a limited number of status updates from multiple users,
// older than the provided timestamp.
func GetStatusesByUserIDs(userIDs []uuid.UUID, cursor time.Time, limit int) ([]Status, error) {
	var statuses []Status
	if err := db.Preload("User").Preload("Media").
		Where("user_id IN (?) AND created_at < ?", userIDs, cursor).
		Order("created_at desc").
		Limit(limit).
		Find(&statuses).Error; err != nil {
		return nil, err
	}
	return statuses, nil
}

// CreateStatus creates a new status update.
func CreateStatus(userID uuid.UUID, text string, media []Media) (*Status, error) {
	// Retrieve the user from the database
	user := User{}
	if err := db.First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}

	newStatus := Status{
		BaseModel: BaseModel{ID: uuid.New()},
		User:      user,
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
