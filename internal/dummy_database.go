package internal

import "github.com/Shikachuu/template-files/pkg"

type DummyDatabase struct {
	templateRecords map[int64]pkg.TemplateRecord
}

var _ pkg.Database = &DummyDatabase{}

func (d *DummyDatabase) GetTemplateById(id int64) (pkg.TemplateRecord, error) {
	templateRecord, ok := d.templateRecords[id]

	if !ok {
		return pkg.TemplateRecord{}, pkg.ErrNotFound
	}

	return templateRecord, nil
}

func (d *DummyDatabase) CreateTemplate(template string, fileExtension string) (pkg.TemplateRecord, error) {
	id := int64(len(d.templateRecords))

	d.templateRecords[id] = pkg.TemplateRecord{
		ID:            id,
		Template:      template,
		FileExtension: fileExtension,
	}

	return d.templateRecords[id], nil
}
