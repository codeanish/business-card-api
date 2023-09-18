package models

type LeetcodeStats struct {
	Data struct {
		MatchedUser struct {
			Profile struct {
				Ranking int `json:"ranking"`
			} `json:"profile"`
		} `json:"matchedUser"`
	} `json:"data"`
}
