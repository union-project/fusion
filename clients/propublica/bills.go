package propublica

import (
	"encoding/json"

	"github.com/union-project/fusion/types"
)

type BillResponse struct {
	Result
	Results []BillResultResponse `json:"results"`
}

type BillResultResponse struct {
	NumberOfResults int          `json:"num_results"`
	Offset          int          `json:"offset"`
	Bills           []types.Bill `json:"bills"`
}

func (c *Client) SearchBills(query string) ([]types.Bill, error) {
	resp, err := c.get("/congress/v1/bills/search.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var billResponse *BillResponse
	if err := json.NewDecoder(resp.Body).Decode(&billResponse); err != nil {
		return nil, err
	}

	bills := []types.Bill{}
	for _, v := range billResponse.Results {
		bills = append(bills, v.Bills...)
	}

	return bills, nil
}
