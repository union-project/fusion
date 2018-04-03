package manager

import (
	"archive/zip"
	"path/filepath"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/union-project/fusion/types"
)

func (m *Manager) ImportBillArchive(archivePath string, maxWorkers int) error {
	r, err := zip.OpenReader(archivePath)
	if err != nil {
		return err
	}
	defer r.Close()

	workerChan := make(chan *zip.File)
	resultChan := make(chan *types.Bill)

	wg := &sync.WaitGroup{}

	// start workers
	for i := 0; i < maxWorkers; i++ {
		go m.importer(i, workerChan, resultChan)
		go m.handler(i, resultChan, wg)
	}

	for _, f := range r.File {
		name := filepath.Base(f.Name)
		if name != "data.json" {
			continue
		}

		logrus.WithFields(logrus.Fields{
			"name": f.Name,
		}).Info("importing file")

		wg.Add(1)

		workerChan <- f
	}

	wg.Wait()

	close(workerChan)
	close(resultChan)

	return err
}

func (m *Manager) importer(id int, c chan *zip.File, r chan *types.Bill) {
	for i := range c {
		f, err := i.Open()
		if err != nil {
			logrus.WithError(err).Error("error importing bill")
			f.Close()
			continue
		}
		bill, err := m.propublicaClient.LoadBill(f)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"id": bill.BillID,
			}).WithError(err).Error("error loading bill")
			f.Close()
			continue
		}

		logrus.WithFields(logrus.Fields{
			"worker":  id,
			"bill_id": bill.BillID,
		}).Debug("loaded bill")

		r <- bill
		f.Close()
	}
}

func (m *Manager) handler(id int, r chan *types.Bill, wg *sync.WaitGroup) {
	for bill := range r {
		if err := m.saveBill(*bill); err != nil {
			logrus.WithFields(logrus.Fields{
				"id": bill.BillID,
			}).WithError(err).Error("error importing bill")
		}

		logrus.WithFields(logrus.Fields{
			"worker": id,
			"id":     bill.BillID,
		}).Debug("imported bill")
		wg.Done()
	}
}
