package router

import (
	"path/filepath"

	"github.com/bwoff11/frens/internal/database"
	"github.com/bwoff11/frens/internal/logger"
	"github.com/bwoff11/frens/internal/response"
	"github.com/bwoff11/frens/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type FilesRepo struct {
	DB  *database.Database
	Srv *service.Service
}

func NewFilesRepo(db *database.Database, srv *service.Service) *FilesRepo {
	return &FilesRepo{
		DB:  db,
		Srv: srv,
	}
}

func (fr *FilesRepo) ConfigureRoutes(rtr fiber.Router) {
	rtr.Post("/", fr.create)
	rtr.Get("/:fileId", fr.getByID)
	rtr.Delete("/:fileId", fr.deleteByID)
}

// create handles the request to create a new file.
// @Summary Create a new file
// @Description Create a new file from the provided form data
// @Tags Files
// @Accept  multipart/form-data
// @Produce  json
// @Param file formData file true "File to upload"
// @Success 200
// @Failure 400
// @Failure 500
// @Security ApiKeyAuth
// @Router /files [post]
func (fr *FilesRepo) create(c *fiber.Ctx) error {

	// Get the file from the request
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	// Validate file exists
	if file == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.CreateErrorResponse(response.ErrInternal))
	}

	return fr.Srv.Files.Create(c, file)
}

// getByID handles the request to get a file by ID.
// @Summary Get a file by ID
// @Description Retrieve a specific file using its ID
// @Tags Files
// @Accept  json
// @Produce  json
// @Param fileId path string true "File ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Security ApiKeyAuth
// @Router /files/{fileId} [get]
func (fr *FilesRepo) getByID(c *fiber.Ctx) error {
	logger.DebugLogRequestRecieved("router", "files", "retrieveFile")

	// Get the file name from the request
	filename := c.Params("fileId")
	if filename == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrFileIDNotProvided))
	}
	logger.Log.Debug().
		Str("package", "router").
		Str("router", "files").
		Str("method", "retrieveFile").
		Str("file_id", filename).
		Msg("Got file ID from request")

	// If extension is provided, remove it
	var fileId string
	if filepath.Ext(filename) != "" {
		fileId = filename[:len(filename)-len(filepath.Ext(filename))]
		logger.Log.Debug().
			Str("package", "router").
			Str("router", "files").
			Str("method", "retrieveFile").
			Str("file_id", fileId).
			Msg("Removed file extension")
	} else {
		fileId = filename
	}

	// Convert the file ID to a UUID
	fileIdUUID, err := uuid.Parse(fileId)
	if err != nil {
		logger.Log.Debug().
			Str("package", "router").
			Str("router", "files").
			Str("method", "retrieveFile").
			Str("file_id", fileId).
			Msg("Failed to convert file ID to UUID")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrFileIDNotUUID))
	}
	logger.Log.Debug().
		Str("package", "router").
		Str("router", "files").
		Str("method", "retrieveFile").
		Str("file_id", fileIdUUID.String()).
		Msg("Converted file ID to UUID")

	// Send the request to the service package
	return fr.Srv.Files.RetrieveByID(c, &fileIdUUID)
}

// deleteByID handles the request to delete a file by ID.
// @Summary Delete a file by ID
// @Description Delete a specific file using its ID
// @Tags Files
// @Accept  json
// @Produce  json
// @Param fileId path string true "File ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Security ApiKeyAuth
// @Router /files/{fileId} [delete]
func (fr *FilesRepo) deleteByID(c *fiber.Ctx) error {
	logger.DebugLogRequestRecieved("router", "files", "deleteFile")

	// Get the file name from the request
	filename := c.Params("fileId")
	logger.Log.Debug().
		Str("package", "router").
		Str("router", "files").
		Str("method", "retrieveFile").
		Str("file_id", filename).
		Msg("Got file ID from request")

	// If extension is provided, remove it
	var fileId string
	if filepath.Ext(filename) != "" {
		fileId = filename[:len(filename)-len(filepath.Ext(filename))]
		logger.Log.Debug().
			Str("package", "router").
			Str("router", "files").
			Str("method", "retrieveFile").
			Str("file_id", fileId).
			Msg("Removed file extension")
	} else {
		fileId = filename
	}

	// Convert the file ID to a UUID
	fileIdUUID, err := uuid.Parse(fileId)
	if err != nil {
		logger.Log.Error().
			Str("package", "router").
			Str("router", "files").
			Str("method", "retrieveFile").
			Str("file_id", fileId).
			Msg("Failed to convert file ID to UUID")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidFileID))
	}
	logger.Log.Debug().
		Str("package", "router").
		Str("router", "files").
		Str("method", "retrieveFile").
		Str("file_id", fileIdUUID.String()).
		Msg("Converted file ID to UUID")

	// Send the request to the service package
	return fr.Srv.Files.DeleteByID(c, &fileIdUUID)
}
