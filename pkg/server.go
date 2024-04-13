package pkg

import (
	"context"

	"github.com/Shikachuu/template-files/proto"
)

type Server struct {
	proto.UnimplementedTemplateServiceServer
	db Database
}

func NewServer(db Database) *Server {
	return &Server{db: db}
}

func (s *Server) GetTemplateById(ctx context.Context, r *proto.GetTemplateRequest) (*proto.TemplateResponse, error) {
	t, err := s.db.GetTemplateById(r.TemplateId)
	if err != nil {
		return nil, err
	}

	tr := &proto.TemplateResponse{
		TemplateId:    t.ID,
		Template:      t.Template,
		FileExtension: &t.FileExtension,
	}
	return tr, nil
}

func (s *Server) CreateTemplate(ctx context.Context, r *proto.TemplateRequest) (*proto.TemplateResponse, error) {
	fe := ""

	if r.FileExtension != nil {
		fe = *r.FileExtension
	}

	t, err := s.db.CreateTemplate(r.Template, fe)
	if err != nil {
		return nil, err
	}

	tr := &proto.TemplateResponse{
		TemplateId:    t.ID,
		Template:      t.Template,
		FileExtension: &t.FileExtension,
	}
	return tr, nil
}
