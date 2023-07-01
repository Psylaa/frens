package service

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/google/uuid"
)

type BookmarkRepo struct{}

func (br *BookmarkRepo) GetByBookmarkID(bookmarkID *uuid.UUID) (int, *response.Response) {
	logger.DebugLogRequestRecieved("service", "bookmark", "GetByBookmarkID")

	// Get bookmark from database
	bookmark, err := db.Bookmarks.GetByID(bookmarkID)
	if err != nil {
		return 0, nil
	}

	resp := response.CreateBookmarkResponse([]*database.Bookmark{bookmark})
	logger.DebugLogRequestCompleted("service", "bookmark", "GetByBookmarkID")
	return 200, resp
}

func (br *BookmarkRepo) GetByPostID(postID *uuid.UUID) (int, *response.Response) {
	logger.DebugLogRequestRecieved("service", "bookmark", "GetByPostID")

	// Get bookmarks from database
	bookmarks, err := db.Bookmarks.GetByPostID(postID)
	if err != nil {
		return 500, nil
	}

	resp := response.CreateBookmarkResponse(bookmarks)
	logger.DebugLogRequestCompleted("service", "bookmark", "GetByPostID")
	return 200, resp
}

func (br *BookmarkRepo) GetCountByPostID(postID *uuid.UUID) (int, *response.Response) {
	logger.DebugLogRequestRecieved("service", "bookmark", "GetCountByPostID")

	// Get bookmark count from database
	count, err := db.Bookmarks.GetCountByPostID(postID)
	if err != nil {
		return 500, nil
	}

	resp := response.CreateCountResponse(count)
	logger.DebugLogRequestCompleted("service", "bookmark", "GetCountByPostID")
	return 200, resp
}

func (br *BookmarkRepo) GetCountByUserID(userID *uuid.UUID) (int, *response.Response) {
	logger.DebugLogRequestRecieved("service", "bookmark", "GetCountByUserID")

	logger.DebugLogRequestCompleted("service", "bookmark", "GetCountByUserID")
	return 0, nil
}

func (br *BookmarkRepo) DeleteByID(bookmarkID *uuid.UUID) (int, *response.Response) {
	logger.DebugLogRequestRecieved("service", "bookmark", "DeleteByID")

	logger.DebugLogRequestCompleted("service", "bookmark", "DeleteByID")
	return 0, nil
}
