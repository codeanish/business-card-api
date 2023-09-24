package models

type BadgeCounts struct {
	Bronze int `json:"bronze"`
	Silver int `json:"silver"`
	Gold   int `json:"gold"`
}

type Item struct {
	BadgeCounts             BadgeCounts `json:"badge_counts"`
	AccountID               int         `json:"account_id"`
	IsEmployee              bool        `json:"is_employee"`
	LastModifiedDate        int64       `json:"last_modified_date"`
	LastAccessDate          int64       `json:"last_access_date"`
	ReputationChangeYear    int         `json:"reputation_change_year"`
	ReputationChangeQuarter int         `json:"reputation_change_quarter"`
	ReputationChangeMonth   int         `json:"reputation_change_month"`
	ReputationChangeWeek    int         `json:"reputation_change_week"`
	ReputationChangeDay     int         `json:"reputation_change_day"`
	Reputation              int         `json:"reputation"`
	CreationDate            int64       `json:"creation_date"`
	UserType                string      `json:"user_type"`
	UserID                  int         `json:"user_id"`
	AcceptRate              int         `json:"accept_rate"`
	WebsiteURL              string      `json:"website_url"`
	Link                    string      `json:"link"`
	ProfileImage            string      `json:"profile_image"`
	DisplayName             string      `json:"display_name"`
}

type StackOverflowResponse struct {
	Items          []Item `json:"items"`
	HasMore        bool   `json:"has_more"`
	QuotaMax       int    `json:"quota_max"`
	QuotaRemaining int    `json:"quota_remaining"`
}
