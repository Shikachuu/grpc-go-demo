package pkg

import (
	"context"

	"github.com/Shikachuu/template-files/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	proto.UnimplementedTemplateServiceServer
	db Database
}

var _ proto.TemplateServiceServer = &Server{}

func NewServer(db Database) *Server {
	return &Server{db: db}
}

func (s *Server) GetTemplateByName(ctx context.Context, r *proto.GetTemplateByNameRequest) (*proto.TemplateResponse, error) {
	t, err := s.db.GetTemplateByName(r.Name)
	if err != nil {
		if err == ErrNotFound {
			return nil, status.Errorf(codes.NotFound, "template with name %s not found", r.Name)
		}

		return nil, status.Errorf(codes.Internal, "failed to get template by name: %v", err)
	}

	tr := &proto.TemplateResponse{
		TemplateId:    t.ID,
		Name:          t.Name,
		Template:      t.Template,
		FileExtension: t.FileExtension,
	}
	return tr, nil
}

func (s *Server) GetTemplateById(ctx context.Context, r *proto.GetTemplateRequest) (*proto.TemplateResponse, error) {
	t, err := s.db.GetTemplateById(r.TemplateId)
	if err != nil {
		if err == ErrNotFound {
			return nil, status.Errorf(codes.NotFound, "template with id %d not found", r.TemplateId)
		}

		return nil, status.Errorf(codes.Internal, "failed to get template by id: %v", err)
	}

	tr := &proto.TemplateResponse{
		TemplateId:    t.ID,
		Name:          t.Name,
		Template:      t.Template,
		FileExtension: t.FileExtension,
	}
	return tr, nil
}

func (s *Server) CreateTemplate(ctx context.Context, r *proto.TemplateRequest) (*proto.TemplateResponse, error) {
	t, err := s.db.CreateTemplate(r.Name, r.Template, r.FileExtension)
	if err != nil {
		if err == ErrDuplicate {
			return nil, status.Errorf(codes.AlreadyExists, "template with name %s already exists", r.Name)
		}
		return nil, status.Errorf(codes.Internal, "failed to create template: %v", err)
	}

	tr := &proto.TemplateResponse{
		TemplateId:    t.ID,
		Name:          t.Name,
		Template:      t.Template,
		FileExtension: t.FileExtension,
	}
	return tr, nil
}
