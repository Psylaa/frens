package router

import (
	"strconv"

	"github.com/bwoff11/frens/service"
	"github.com/gofiber/fiber/v2"
)

type LikesRepo struct {
	Service *service.LikeService
}

func (lr *LikesRepo) addPrivateRoutes(rtr fiber.Router) {
	grp := rtr.Group("/likes")
	grp.Post("/:postID", lr.likePost)
	grp.Delete("/:postID", lr.unlikePost)
}

func (lr *LikesRepo) likePost(c *fiber.Ctx) error {
	postID, err := lr.parseID(c.Params("postID"))
	if err != nil {
		return err
	}

	return lr.Service.LikePost(c, postID)
}

func (lr *LikesRepo) unlikePost(c *fiber.Ctx) error {
	postID, err := lr.parseID(c.Params("postID"))
	if err != nil {
		return err
	}

	return lr.Service.UnlikePost(c, postID)
}

func (lr *LikesRepo) parseID(idStr string) (uint32, error) {
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(id), nil
}
