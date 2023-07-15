package models

type FeedResponse struct {
	Links    FeedLinks  `json:"links"`
	Data     []PostData `json:"data"`
	Included []UserData `json:"included"`
}

type FeedLinks struct {
	Self string `json:"self"`
}

func CreateFeedResponse(data []PostData, included []UserData) *FeedResponse {
	// Create links
	links := FeedLinks{
		Self: "/feed",
	}

	// Create response
	return &FeedResponse{
		Links:    links,
		Data:     data,
		Included: included,
	}
}
