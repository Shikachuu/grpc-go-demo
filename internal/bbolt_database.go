package internal

import (
	"encoding/json"

	"github.com/Shikachuu/template-files/pkg"
	bolt "go.etcd.io/bbolt"
)

const templatesBucket = "templates"

type BBoltDatabase struct {
	db *bolt.DB
}

var _ pkg.Database = &BBoltDatabase{}

func NewBBoltDatabase(db *bolt.DB) *BBoltDatabase {
	return &BBoltDatabase{db: db}
}

// CreateTemplate implements pkg.Database.
func (b *BBoltDatabase) CreateTemplate(name string, template string, fileExtension *string) (pkg.TemplateRecord, error) {
	var record pkg.TemplateRecord

	err := b.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(templatesBucket))
		if err != nil {
			return err
		}

		id, err := bucket.NextSequence()
		if err != nil {
			return err
		}

		record = pkg.TemplateRecord{
			ID:            int64(id),
			Name:          name,
			Template:      template,
			FileExtension: fileExtension,
		}

		bb, err := json.Marshal(record)
		if err != nil {
			return err
		}

		err = bucket.Put([]byte(name), bb)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return record, err
	}

	return record, nil
}

// GetTemplateById implements pkg.Database.
func (b *BBoltDatabase) GetTemplateById(id int64) (pkg.TemplateRecord, error) {
	var record pkg.TemplateRecord

	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(templatesBucket))
		if bucket == nil {
			return pkg.ErrNotFound
		}

		err := bucket.ForEach(func(k, v []byte) error {
			var t pkg.TemplateRecord
			err := json.Unmarshal(v, &t)
			if err != nil {
				return err
			}

			if t.ID == id {
				record = t
				return nil
			}

			return pkg.ErrNotFound
		})

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return record, err
	}

	return record, nil
}

// GetTemplateByName implements pkg.Database.
func (b *BBoltDatabase) GetTemplateByName(name string) (pkg.TemplateRecord, error) {
	var record pkg.TemplateRecord

	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(templatesBucket))
		if bucket == nil {
			return pkg.ErrNotFound
		}

		t := bucket.Get([]byte(name))
		if t == nil {
			return pkg.ErrNotFound
		}

		err := json.Unmarshal(t, &record)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return record, err
	}

	return record, nil
}
