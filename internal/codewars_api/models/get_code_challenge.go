package models

import "time"

type GetCodeChallengeResponse struct {
	Id             string                           `json:"id"`
	Name           string                           `json:"name"`
	Slug           string                           `json:"slug"`
	Url            string                           `json:"url"`
	Category       string                           `json:"category"`
	Description    string                           `json:"description"`
	Tags           []string                         `json:"tags"`
	Languages      []string                         `json:"languages"`
	Rank           Rank                             `json:"rank"`
	CreatedBy      GetCodeChallengeResponseActionBy `json:"createdBy"`
	ApprovedBy     GetCodeChallengeResponseActionBy `json:"approvedBy"`
	TotalAttempts  int                              `json:"totalAttempts"`
	TotalCompleted int                              `json:"totalCompleted"`
	TotalStars     int                              `json:"totalStars"`
	VoteScore      int                              `json:"voteScore"`
	PublishedAt    time.Time                        `json:"publishedAt"`
	ApprovedAt     time.Time                        `json:"approvedAt"`
}

type GetCodeChallengeResponseActionBy struct {
	Username string `json:"username"`
	Url      string `json:"url"`
}
