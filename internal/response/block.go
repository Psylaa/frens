package response

import "github.com/bwoff11/frens/internal/database"

type BlocksResponse struct{}

func CreateBlocksResponse(blocks []database.Block) *BlocksResponse {
	return nil
}
