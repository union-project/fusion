package types

type Result struct {
	Status    string `json:"status"`
	Copyright string `json:"copyright"`
}

type MemberResponse struct {
	Result
	Results []ResultResponse `json:"results"`
}

type ResultResponse struct {
	Congress        string   `json:"congress"`
	Chamber         string   `json:"chamber"`
	NumberOfResults int      `json:"num_results"`
	Offset          int      `json:"offset"`
	Members         []Member `json:"members"`
}

type Member struct {
	ID             int    `json:"_id" gorm:"column:id;PRIMARY_KEY"`
	MemberID       string `json:"id" gorm:"column:member_id"`
	Title          string `json:"title,omitempty"`
	ShortTitle     string `json:"short_title,omitempty"`
	APIUrl         string `json:"api_url,omitempty"`
	FirstName      string `json:"first_name,omitempty"`
	MiddleName     string `json:"middle_name,omitempty"`
	LastName       string `json:"last_name,omitempty"`
	Suffix         string `json:"suffix,omitempty"`
	DateOfBirth    string `json:"date_of_birth,omitempty"`
	Party          string `json:"party,omitempty"`
	LeadershipRole string `json:"leadership_role,omitempty"`
	Twitter        string `json:"twitter_account,omitempty"`
	Facebook       string `json:"facebook_account,omitempty"`
	Youtube        string `json:"youtube_account,omitempty"`
	GovtrackID     string `json:"govtrack_id,omitempty"`
	CSPANID        string `json:"cspan_id,omitempty"`
	VotesmartID    string `json:"votesmart_id,omitempty"`

	State                    string  `json:"state,omitempty"`
	VotesWithPartyPercentage float64 `json:"votes_with_party_pct,omitempty"`
}
