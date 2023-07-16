package like

import "github.com/bwoff11/frens/pkg/database"

type Service struct {
	Database *database.Database
}

func New(db *database.Database) *Service {
	return &Service{
		Database: db,
	}
}
