package pkg_test

import (
	"context"
	"log"
	"net"
	"os"
	"testing"

	"github.com/Shikachuu/template-files/internal"
	"github.com/Shikachuu/template-files/pkg"
	"github.com/Shikachuu/template-files/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	pb "google.golang.org/protobuf/proto"
)

var (
	lis            *bufconn.Listener
	dummyExtension                        = "txt"
	db             internal.DummyDatabase = internal.DummyDatabase{TemplateRecords: map[int64]pkg.TemplateRecord{
		1: {ID: 1, Name: "test", Template: "template", FileExtension: &dummyExtension},
	}}
)

func TestMain(m *testing.M) {
	lis = bufconn.Listen(1024 * 1024)
	defer lis.Close()

	s := grpc.NewServer()
	defer s.Stop()

	proto.RegisterTemplateServiceServer(s, pkg.NewServer(&db))
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("server exited with error: %v", err)
		}
	}()
	os.Exit(m.Run())
}

func newClient() proto.TemplateServiceClient {
	ctx := context.Background()
	do := grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
		return lis.Dial()
	})
	conn, err := grpc.DialContext(ctx, "bufnet", do, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial bufnet: %v", err)
	}
	return proto.NewTemplateServiceClient(conn)
}

func Test_GetTemplateById(t *testing.T) {
	client := newClient()
	testCases := []struct {
		input           *proto.GetTemplateRequest
		desc            string
		expected        *proto.TemplateResponse
		expectedErrCode codes.Code
	}{
		{
			desc:     "get template by id positive",
			input:    &proto.GetTemplateRequest{TemplateId: 1},
			expected: &proto.TemplateResponse{TemplateId: 1, Name: "test", Template: "template", FileExtension: &dummyExtension},
		},
		{
			desc:            "get template by invalid id",
			input:           &proto.GetTemplateRequest{TemplateId: 4},
			expectedErrCode: codes.NotFound,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r, err := client.GetTemplateById(context.Background(), tC.input)
			if err != nil {
				if tC.expectedErrCode == codes.OK {
					t.Fatalf("failed to get template by id: %v", err)
				}

				st, ok := status.FromError(err)
				if !ok || st.Code() != tC.expectedErrCode {
					t.Fatalf("expected error %v, got %v", tC.expected, r)
				}
			}

			if !pb.Equal(r, tC.expected) {
				t.Fatalf("expected %v, got %v", tC.expected, r)
			}
		})
	}
}

func Test_GetTemplateByName(t *testing.T) {
	client := newClient()
	testCases := []struct {
		input           *proto.GetTemplateByNameRequest
		desc            string
		expected        *proto.TemplateResponse
		expectedErrCode codes.Code
	}{
		{
			desc:     "get template by name positive",
			input:    &proto.GetTemplateByNameRequest{Name: "test"},
			expected: &proto.TemplateResponse{TemplateId: 1, Name: "test", Template: "template", FileExtension: &dummyExtension},
		},
		{
			desc:            "get template by invalid name",
			input:           &proto.GetTemplateByNameRequest{Name: "invalid"},
			expectedErrCode: codes.NotFound,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r, err := client.GetTemplateByName(context.Background(), tC.input)
			if err != nil {
				if tC.expectedErrCode == codes.OK {
					t.Fatalf("failed to create template: %v", err)
				}

				st, ok := status.FromError(err)
				if !ok || st.Code() != tC.expectedErrCode {
					t.Fatalf("expected error %v, got %v", tC.expected, r)
				}
			}

			if !pb.Equal(r, tC.expected) {
				t.Fatalf("expected %v, got %v", tC.expected, r)
			}
		})
	}
}

func Test_CreateTemplate(t *testing.T) {
	client := newClient()
	testCases := []struct {
		input           *proto.TemplateRequest
		desc            string
		expected        *proto.TemplateResponse
		expectedErrCode codes.Code
	}{
		{
			desc:     "create template positive",
			input:    &proto.TemplateRequest{Template: "new template", FileExtension: &dummyExtension},
			expected: &proto.TemplateResponse{TemplateId: 2, Template: "new template", FileExtension: &dummyExtension},
		},
		{
			desc:            "duplicate name",
			input:           &proto.TemplateRequest{Template: "test", FileExtension: &dummyExtension},
			expectedErrCode: codes.AlreadyExists,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			r, err := client.CreateTemplate(context.Background(), tC.input)
			if err != nil {
				if tC.expectedErrCode == codes.OK {
					t.Fatalf("failed to create template: %v", err)
				}

				st, ok := status.FromError(err)
				if !ok || st.Code() != tC.expectedErrCode {
					t.Fatalf("expected error %v, got %v", tC.expected, r)
				}
			}

			if !pb.Equal(r, tC.expected) {
				t.Fatalf("expected %v, got %v", tC.expected, r)
			}
		})
	}
}
