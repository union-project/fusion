package manager

import "github.com/union-project/fusion/types"

func (m *Manager) Members() ([]types.Member, error) {
	var members []types.Member
	if err := m.db.Find(&members).Error; err != nil {
		return nil, err
	}

	return members, nil
}
