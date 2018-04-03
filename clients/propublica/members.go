package propublica

import (
	"encoding/json"

	"github.com/union-project/fusion/types"
)

type Result struct {
	Status    string `json:"status"`
	Copyright string `json:"copyright"`
}

type MemberResponse struct {
	Result
	Results []MemberResultResponse `json:"results"`
}

type MemberResultResponse struct {
	Congress        string         `json:"congress"`
	Chamber         string         `json:"chamber"`
	NumberOfResults int            `json:"num_results"`
	Offset          int            `json:"offset"`
	Members         []types.Member `json:"members"`
}

func (c *Client) Members() ([]types.Member, error) {
	resp, err := c.get("/congress/v1/115/senate/members.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var memberResponse *MemberResponse
	if err := json.NewDecoder(resp.Body).Decode(&memberResponse); err != nil {
		return nil, err
	}

	members := []types.Member{}
	for _, v := range memberResponse.Results {
		members = append(members, v.Members...)
	}

	return members, nil
}
