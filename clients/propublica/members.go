package propublica

import (
	"encoding/json"

	"github.com/union-project/fusion/types"
)

func (c *Client) Members() ([]types.Member, error) {
	resp, err := c.get("/congress/v1/115/senate/members.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var memberResponse *types.MemberResponse
	if err := json.NewDecoder(resp.Body).Decode(&memberResponse); err != nil {
		return nil, err
	}

	members := []types.Member{}
	for _, v := range memberResponse.Results {
		members = append(members, v.Members...)
	}

	return members, nil
}
