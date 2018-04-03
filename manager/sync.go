package manager

import (
	"github.com/sirupsen/logrus"
	"github.com/union-project/fusion/types"
)

func (m *Manager) Sync() error {
	if err := m.syncMembers(); err != nil {
		return err
	}

	if err := m.syncBills(); err != nil {
		return err
	}

	return nil
}

func (m *Manager) syncMembers() error {
	members, err := m.propublicaClient.Members()
	if err != nil {
		return err
	}

	for _, member := range members {
		if err := m.db.FirstOrCreate(&member, types.Member{MemberID: member.MemberID}).Error; err != nil {
			logrus.WithFields(logrus.Fields{
				"member": member.MemberID,
			}).WithError(err).Error("unable to update member")
			continue
		}
		logrus.WithFields(logrus.Fields{
			"id":         member.MemberID,
			"title":      member.Title,
			"first_name": member.FirstName,
			"last_name":  member.LastName,
			"state":      member.State,
		}).Info("updating member")
	}

	return nil
}

func (m *Manager) syncBills() error {
	bills, err := m.propublicaClient.SearchBills("")
	if err != nil {
		return err
	}

	for _, bill := range bills {
		if err := m.saveBill(bill); err != nil {
			logrus.WithFields(logrus.Fields{
				"id": bill.BillID,
			}).WithError(err).Error("unable to update bill")
			continue
		}
		logrus.WithFields(logrus.Fields{
			"id":    bill.BillID,
			"title": bill.Title,
		}).Info("updating bill")
	}

	return nil
}
