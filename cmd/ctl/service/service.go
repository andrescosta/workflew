package service

import (
	"context"

	"github.com/andrescosta/goico/pkg/database"
	"github.com/andrescosta/goico/pkg/env"
	"github.com/andrescosta/goico/pkg/service"
	"github.com/andrescosta/goico/pkg/service/grpc"
	pb "github.com/andrescosta/jobico/internal/api/types"
	"github.com/andrescosta/jobico/internal/ctl/server"
)

const name = "ctl"

type Service struct {
	Listener service.GrpcListener
	DBOption *database.Option
	Dialer   service.GrpcDialer
}

func (s Service) Start(ctx context.Context) error {
	l := s.Listener
	if l == nil {
		l = service.DefaultGrpcListener
	}
	o := s.DBOption
	if o == nil {
		o = &database.Option{}
	}
	svc, err := grpc.New(
		grpc.WithName(name),
		grpc.WithListener(l),
		grpc.WithContext(ctx),
		grpc.WithServiceDesc(&pb.Control_ServiceDesc),
		grpc.WithNewServiceFn(func(ctx context.Context) (any, error) {
			dbFileName := env.String("ctl.dbname", "db.db")
			return server.New(ctx, dbFileName, *o)
		}),
	)
	if err != nil {
		return err
	}
	defer svc.Dispose()
	if err = svc.Serve(); err != nil {
		return err
	}
	return nil
}

func (s Service) Addr() *string {
	return env.StringOrNil(name + ".addr")
}

func (s Service) Kind() service.Kind {
	return grpc.Kind
}

func (s Service) CheckHealth(ctx context.Context) error {
	d := s.Dialer
	if d == nil {
		d = service.DefaultGrpcDialer
	}
	cli, err := grpc.NewHelthCheckClient(ctx, *s.Addr(), d)
	if err != nil {
		return err
	}
	defer cli.Close()
	return cli.CheckOk(ctx, name)
}
