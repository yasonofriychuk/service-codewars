package models

import "time"

type ListCompletedChallengesResponse struct {
	TotalPages int                                   `json:"totalPages"`
	TotalItems int                                   `json:"totalItems"`
	Data       []ListCompletedChallengesResponseData `json:"data"`
}

type ListCompletedChallengesResponseData struct {
	Id                 string    `json:"id"`
	Name               string    `json:"name"`
	Slug               string    `json:"slug"`
	CompletedAt        time.Time `json:"completedAt"`
	CompletedLanguages []string  `json:"completedLanguages"`
}
