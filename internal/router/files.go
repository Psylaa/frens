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
	rtr.Delete("/:fileID", fr.deleteByID)
}

// @Summary Search Files
// @Description Search for files with query parameters. If no query parameters are provided, all files will be returned.
// @Tags Files
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400
// @Failure 401
// @Failure 500
// @Security ApiKeyAuth
// @Router /files [get]
func (fr *FilesRepo) search(c *fiber.Ctx) error {
	return nil
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

// deleteByID handles the request to delete a file by ID.
// @Summary Delete a file by ID
// @Description Delete a specific file using its ID
// @Tags Files
// @Accept  json
// @Produce  json
// @Param fileID path string true "File ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Security ApiKeyAuth
// @Router /files/{fileID} [delete]
func (fr *FilesRepo) deleteByID(c *fiber.Ctx) error {
	logger.DebugLogRequestReceived("router", "files", "deleteFile")

	// Get the file name from the request
	filename := c.Params("fileID")
	logger.Log.Debug().
		Str("package", "router").
		Str("router", "files").
		Str("method", "retrieveFile").
		Str("file_id", filename).
		Msg("Got file ID from request")

	// If extension is provided, remove it
	var fileID string
	if filepath.Ext(filename) != "" {
		fileID = filename[:len(filename)-len(filepath.Ext(filename))]
		logger.Log.Debug().
			Str("package", "router").
			Str("router", "files").
			Str("method", "retrieveFile").
			Str("file_id", fileID).
			Msg("Removed file extension")
	} else {
		fileID = filename
	}

	// Convert the file ID to a UUID
	fileIDUUID, err := uuid.Parse(fileID)
	if err != nil {
		logger.Log.Error().
			Str("package", "router").
			Str("router", "files").
			Str("method", "retrieveFile").
			Str("file_id", fileID).
			Msg("Failed to convert file ID to UUID")
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidFileID))
	}
	logger.Log.Debug().
		Str("package", "router").
		Str("router", "files").
		Str("method", "retrieveFile").
		Str("file_id", fileIDUUID.String()).
		Msg("Converted file ID to UUID")

	// Send the request to the service package
	return fr.Srv.Files.DeleteByID(c, &fileIDUUID)
}
