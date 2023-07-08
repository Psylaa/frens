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

// @Summary Retrieve User's Files
// @Description Retrieves a list of files uploaded by the authenticated user.
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

// @Summary Upload File
// @Description Uploads a new file and assigns it to the authenticated user.
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

// @Summary Delete File
// @Description Deletes the specified file from the authenticated user's uploaded files.
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

	// If extension is provided, remove it
	var fileID string
	if filepath.Ext(filename) != "" {
		fileID = filename[:len(filename)-len(filepath.Ext(filename))]
	} else {
		fileID = filename
	}

	// Convert the file ID to a UUID
	fileIDUUID, err := uuid.Parse(fileID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.CreateErrorResponse(response.ErrInvalidFileID))
	}

	// Send the request to the service package
	return fr.Srv.Files.DeleteByID(c, &fileIDUUID)
}
