package types

import "github.com/jinzhu/gorm"

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
	ID           int     `json:"_id" gorm:"column:id;primary_key;auto_increment"`
	BillID       string  `json:"bill_id,omitempty"`
	BillType     string  `json:"bill_type,omitempty"`
	ByRequest    bool    `json:"by_request,omitempty"`
	Number       string  `json:"number,omitempty"`
	Title        string  `json:"official_title,omitempty"`
	Sponsor      Sponsor `json:"sponsor,omitempty" gorm:"foreignkey:SponsorID"`
	IntroducedAt string  `json:"introduced_at,omitempty"`
	Status       string  `json:"status,omitempty"`
	StatusAt     string  `json:"status_at,omitempty"`
	Summary      Summary `json:"summary,omitempty" gorm:"foreignkey:SummaryID"`
	UpdatedAt    string  `json:"updated_at,omitempty"`
}

type Sponsor struct {
	gorm.Model
	Name  string `json:"name,omitempty"`
	State string `json:"state,omitempty"`
	Title string `json:"title,omitempty"`
	Type  string `json:"type,omitempty"`
}

type Summary struct {
	gorm.Model
	As   string `json:"as,omitempty"`
	Date string `json:"date,omitempty"`
	Text string `json:"text,omitempty"`
}
