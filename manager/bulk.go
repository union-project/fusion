package manager

import (
	"archive/zip"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func (m *Manager) ImportBillArchive(archivePath string) error {
	r, err := zip.OpenReader(archivePath)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		name := filepath.Base(f.Name)
		if name != "data.json" {
			continue
		}

		c, err := f.Open()
		if err != nil {
			return err
		}

		bill, err := m.propublicaClient.LoadBill(c)
		if err != nil {
			return err
		}

		logrus.WithFields(logrus.Fields{
			"title": bill.Title,
		}).Debug("loading bill")
		if err := m.saveBill(*bill); err != nil {
			return err
		}

		c.Close()
	}

	return err
}
