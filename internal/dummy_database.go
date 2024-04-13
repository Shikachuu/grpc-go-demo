package internal

import (
	"github.com/Shikachuu/template-files/pkg"
)

type DummyDatabase struct {
	TemplateRecords map[int64]pkg.TemplateRecord
}

var _ pkg.Database = &DummyDatabase{}

func (d *DummyDatabase) GetTemplateById(id int64) (pkg.TemplateRecord, error) {
	templateRecord, ok := d.TemplateRecords[id]

	if !ok {
		return pkg.TemplateRecord{}, pkg.ErrNotFound
	}

	return templateRecord, nil
}

func (d *DummyDatabase) CreateTemplate(name string, template string, fileExtension *string) (pkg.TemplateRecord, error) {
	for _, tr := range d.TemplateRecords {
		if tr.Name == name {
			return pkg.TemplateRecord{}, pkg.ErrDuplicate
		}
	}

	id := int64(len(d.TemplateRecords) + 1)

	d.TemplateRecords[id] = pkg.TemplateRecord{
		ID:            id,
		Name:          name,
		Template:      template,
		FileExtension: fileExtension,
	}

	return d.TemplateRecords[id], nil
}
