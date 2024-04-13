package pkg

import "errors"

type TemplateRecord struct {
	ID            int64
	Template      string
	FileExtension string
}

var ErrNotFound = errors.New("template not found")

type Database interface {
	GetTemplateById(id int64) (TemplateRecord, error)
	CreateTemplate(template string, fileExtension string) (TemplateRecord, error)
}
