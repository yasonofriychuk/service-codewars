package models

type GetUserResponse struct {
	Username            string                        `json:"username"`
	Name                string                        `json:"name"`
	Honor               int                           `json:"honor"`
	Clan                string                        `json:"clan"`
	LeaderboardPosition int                           `json:"leaderboardPosition"`
	Skills              []string                      `json:"skills"`
	Ranks               GetUserResponseRanks          `json:"ranks"`
	CodeChallenges      GetUserResponseCodeChallenges `json:"codeChallenges"`
}

type GetUserResponseRanks struct {
	Overall   Rank            `json:"overall"`
	Languages map[string]Rank `json:"languages"`
}

type GetUserResponseCodeChallenges struct {
	TotalAuthored  int `json:"totalAuthored"`
	TotalCompleted int `json:"totalCompleted"`
}
