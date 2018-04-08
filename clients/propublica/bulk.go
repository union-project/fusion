package propublica

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/union-project/fusion/types"
)

type bulkBill struct {
	BillID       string  `json:"bill_id,omitempty"`
	BillType     string  `json:"bill_type,omitempty"`
	Title        string  `json:"official_title,omitempty"`
	Number       string  `json:"number,omitempty"`
	IntroducedAt string  `json:"introduced_at,omitempy"`
	Summary      summary `json:"summary,omitempty"`
	Sponsor      sponsor `json:"sponsor,omitempty"`
	UpdatedAt    string  `json:"updated_at,omitempty"`
}
type sponsor struct {
	Name  string `json:"name,omitempty"`
	State string `json:"state,omitempty"`
	Title string `json:"title,omitempty"`
	Type  string `json:"type,omitempty"`
}
type summary struct {
	As   string `json:"as,omitempty"`
	Date string `json:"date,omitempty"`
	Text string `json:"text,omitempty"`
}

// LoadBill loads bill content into the datastore from a reader
func (c *Client) LoadBill(content io.Reader) (*types.Bill, error) {
	data, err := ioutil.ReadAll(content)
	if err != nil {
		return nil, err
	}

	// detect if bulk or search; bulk has a richer summary field
	var tmp map[string]interface{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return nil, err
	}

	if _, ok := tmp["summary"].(string); ok {
		var bill types.Bill
		if err := json.Unmarshal(data, &bill); err != nil {
			return nil, err
		}

		return &bill, nil
	}

	var b bulkBill
	if err := json.Unmarshal(data, &b); err != nil {
		return nil, err
	}

	return &types.Bill{
		BillID:         b.BillID,
		BillType:       b.BillType,
		Number:         b.Number,
		Title:          b.Title,
		Sponsor:        b.Sponsor.Name,
		Introduced:     b.IntroducedAt,
		Summary:        b.Summary.Text,
		LastActionDate: b.UpdatedAt,
	}, nil
}
