package types

type Member struct {
	ID             int    `json:"_id" gorm:"column:id;primary_key;auto_increment"`
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

type Bill struct {
	ID             int    `json:"_id" gorm:"column:id;primary_key;auto_increment"`
	BillID         string `json:"bill_id,omitempty"`
	BillType       string `json:"bill_type,omitempty"`
	Bill           string `json:"bill,omitempty"`
	Number         string `json:"number,omitempty"`
	Title          string `json:"title,omitempty"`
	ShortTitle     string `json:"short_title,omitempty"`
	Sponsor        string `json:"sponsor,omitempty"`
	Introduced     string `json:"introduced_date,omitempty"`
	Active         bool   `json:"active,omitempty"`
	LastVote       string `json:"last_vote,omitempty"`
	HousePassage   string `json:"house_passage,omitempty"`
	SenatePassage  string `json:"senate_passage,omitempty"`
	Summary        string `json:"summary,omitempty"`
	LastAction     string `json:"latest_major_action,omitempty"`
	LastActionDate string `json:"latest_major_action_date,omitempty"`
}
