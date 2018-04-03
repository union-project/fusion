package propublica

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/union-project/fusion/types"
)

// LoadBill loads bill content into the datastore from a reader
func (c *Client) LoadBill(content io.Reader) (*types.Bill, error) {
	data, err := ioutil.ReadAll(content)
	if err != nil {
		return nil, err
	}

	var bill types.Bill
	if err := json.Unmarshal(data, &bill); err != nil {
		return nil, err
	}

	return &bill, nil
}
