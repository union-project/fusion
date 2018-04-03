package manager

import (
	"github.com/union-project/fusion"
	"github.com/union-project/fusion/types"
)

func (m *Manager) Members(opts ...fusion.ListMemberOpt) ([]types.Member, error) {
	var members []types.Member
	db := m.db
	for _, o := range opts {
		d, err := o(db)
		if err != nil {
			return nil, err
		}
		db = d
	}

	if err := db.Find(&members).Error; err != nil {
		return nil, err
	}

	return members, nil
}
