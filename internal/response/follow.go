package response

import (
	"github.com/bwoff11/frens/internal/database"
	"github.com/google/uuid"
)

type FollowResp struct {
	Links    FollowResp_Links      `json:"links,omitempty"`
	Data     []FollowResp_Data     `json:"data,omitempty"`
	Included []FollowResp_Included `json:"included,omitempty"`
}

type FollowResp_Links struct {
	Self   string `json:"self"`
	Source string `json:"source"`
	Target string `json:"target"`
}

type FollowResp_Data struct {
	Type       string                    `json:"type"`
	ID         string                    `json:"id"`
	Attributes FollowResp_DataAttributes `json:"attributes"`
}

type FollowResp_DataAttributes struct {
	SourceID uuid.UUID `json:"sourceId"`
	TargetID uuid.UUID `json:"targetId"`
}

type FollowResp_Included struct {
	// For future use
}

func GenerateFollowResponse(follow *database.Follow) *FollowResp {
	return &FollowResp{
		Links: FollowResp_Links{
			Self:   "/follows",
			Source: "/users/" + follow.SourceID.String(),
			Target: "/users/" + follow.TargetID.String(),
		},
		Data: []FollowResp_Data{
			{
				Type: "follow",
				ID:   "1",
				Attributes: FollowResp_DataAttributes{
					SourceID: follow.SourceID,
					TargetID: follow.TargetID,
				},
			},
		},
	}
}

func GenerateFollowsResponse(follows []*database.Follow) *FollowResp {
	return &FollowResp{
		Links: FollowResp_Links{
			Self: "/follows",
		},
		Data: []FollowResp_Data{
			{
				Type:       "follow",
				ID:         "1",
				Attributes: FollowResp_DataAttributes{
					// For future use
				},
			},
			{
				Type:       "follow",
				ID:         "2",
				Attributes: FollowResp_DataAttributes{
					// For future use
				},
			},
		},
	}
}
