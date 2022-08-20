package runner

import (
	"context"
	"fmt"
	awesomev1c "github.com/abitofhelp/awesome/client/client_go"
	"github.com/abitofhelp/awesome2/config"
	awesome2v1 "github.com/abitofhelp/awesome2/gen/go/awesome2/v1"
	messagesv1 "github.com/abitofhelp/awesome2/gen/go/awesome2/v1/messages"
	"github.com/abitofhelp/awesome2/server/runner/mapper"
	logger "github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
)

func Run(appcfg *config.AppConfig) error {

	server := grpc.NewServer()

	svc, err := NewService()
	if err != nil {
		panic(err)
	}

	awesome2v1.RegisterAwesome2ServiceServer(server, svc)
	reflection.Register(server)

	con, err := net.Listen("tcp", fmt.Sprintf("%s:%d", appcfg.Runtime.Host, appcfg.Runtime.GrpcPort))
	if err != nil {
		panic(err)
	}

	logger.Printf(">>>> Starting gRPC project service on %s...\n", con.Addr().String())
	return server.Serve(con)
}

type Server struct {
	awesome2v1.Awesome2ServiceServer
}

func (s Server) FindReportByOwnerName(
	ctx context.Context,
	in *awesome2v1.FindReportByOwnerNameRequest) (*awesome2v1.FindReportByOwnerNameResponse, error) {
	return s.findReportByOwnerName(ctx, in)
}

func (s Server) findReportByOwnerName(
	ctx context.Context,
	in *awesome2v1.FindReportByOwnerNameRequest,
	opts ...grpc.CallOption) (*awesome2v1.FindReportByOwnerNameResponse, error) {

	if client, err := awesomev1c.NewAwesomeServiceClient("localhost", 20580); err == nil {
		if report, err := client.FindReportByPetName(context.Background(), "Lassie"); err == nil {
			fmt.Printf("\nFound a report for '%s'\n", report.Report.Pet.Name)

			if at, err := mapper.KeyToAccessTier(report.Report.AccessTier); err == nil {
				if p, err := mapper.KeyToPrivacy(report.Privacy); err == nil {
					return &awesome2v1.FindReportByOwnerNameResponse{
						Report: &messagesv1.Report{
							AccessTier:   at,
							Archived:     report.Report.Archived,
							GeneratedUtc: timestamppb.New(report.Report.GeneratedUtc),
							Pet: &messagesv1.Pet{
								BirthdayUtc: timestamppb.New(report.Report.Pet.BirthdayUtc),
								Name:        report.Report.Pet.Name,
							},
							Title: report.Report.Title,
							Uri:   report.Report.URI,
						},
						Privacy: p,
					}, nil
				} else {
					return nil, err
				}
			} else {
				panic(err)
			}
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}
}

func NewService() (*Server, error) {
	return &Server{}, nil
}
