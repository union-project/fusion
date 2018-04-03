package manager

import (
	"github.com/sirupsen/logrus"
	"github.com/union-project/fusion/clients/propublica"
	"github.com/union-project/fusion/types"
)

func (m *Manager) Sync() error {
	client, err := propublica.NewClient(&propublica.Config{
		APIKey: m.Config.ProPublicaAPIKey,
	})
	if err != nil {
		return err
	}

	members, err := client.Members()
	if err != nil {
		return err
	}

	for _, member := range members {
		if err := m.createOrUpdateMember(&member); err != nil {
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

func (m *Manager) createOrUpdateMember(member *types.Member) error {
	//var mbr *types.Member
	//if err := m.db.Find(&mbr, member.MemberID).Error; err != nil {
	//	logrus.Errorf("unable to find member %s", member.MemberID)
	//	return err
	//}

	//if mbr != nil {
	//	return m.db.Model(mbr).Where("MemberID = ?", member.MemberID).Update(member).Error
	//}

	return m.db.FirstOrCreate(member, types.Member{MemberID: member.MemberID}).Error
}
