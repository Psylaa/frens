package user

import "github.com/bwoff11/frens/pkg/database"

type Service struct {
	Database *database.Database
}

func New(db *database.Database) *Service {
	return &Service{
		Database: db,
	}
}

func (s *Service) Register(username, email, password string) error {
	newUser := &database.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	if err := s.Database.Conn.Create(newUser).Error; err != nil {
		return err
	}

	return nil
}
