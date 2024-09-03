package models

type ListAuthoredChallengesResponse struct {
	Data []ListAuthoredChallengesResponseDataItem `json:"data"`
}

type ListAuthoredChallengesResponseDataItem struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Rank        int      `json:"rank"`
	RankName    string   `json:"rankName"`
	Tags        []string `json:"tags"`
	Languages   []string `json:"languages"`
}
