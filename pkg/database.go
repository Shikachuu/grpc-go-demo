package pkg

import "errors"

type TemplateRecord struct {
	ID            int64
	Name          string
	Template      string
	FileExtension *string
}

var ErrNotFound = errors.New("template not found")
var ErrDuplicate = errors.New("template with the same name already exists")

type Database interface {
	GetTemplateById(id int64) (TemplateRecord, error)
	GetTemplateByName(name string) (TemplateRecord, error)
	CreateTemplate(name string, template string, fileExtension *string) (TemplateRecord, error)
}
