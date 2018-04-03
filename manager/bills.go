package manager

import (
	"github.com/union-project/fusion"
	"github.com/union-project/fusion/types"
)

func (m *Manager) Bills(opts ...fusion.ListBillOpt) ([]types.Bill, error) {
	var bills []types.Bill
	db := m.db
	for _, o := range opts {
		d, err := o(db)
		if err != nil {
			return nil, err
		}
		db = d
	}

	if err := db.Find(&bills).Error; err != nil {
		return nil, err
	}

	return bills, nil
}
