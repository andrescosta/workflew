package service

import (
	"context"
	"errors"

	"github.com/andrescosta/goico/pkg/env"
	"github.com/andrescosta/goico/pkg/service"
	"github.com/andrescosta/goico/pkg/service/http"
	"github.com/andrescosta/jobico/internal/listener"
)

const name = "listener"

type Service struct {
	Dialer        service.GrpcDialer
	Listener      service.HTTPListener
	ListenerCache service.HTTPListener
	ClientBuilder service.HTTPClient
}

func (s Service) Start(ctx context.Context) (err error) {
	d := s.Dialer
	if d == nil {
		d = service.DefaultGrpcDialer
	}
	l := s.Listener
	if l == nil {
		l = service.DefaultHTTPListener
	}
	lc := s.ListenerCache
	if lc == nil {
		lc = service.DefaultGrpcListener
	}
	c, err := listener.New(ctx, d, lc)
	if err != nil {
		return
	}
	defer func() {
		err = errors.Join(err, c.Close())
	}()
	svc, err := http.New(
		http.WithListener[*http.ServiceOptions](l),
		http.WithContext(ctx),
		http.WithName(name),
		http.WithHealthCheck[*http.ServiceOptions](func(ctx context.Context) (map[string]string, error) {
			return make(map[string]string), nil
		}),
		http.WithInitRoutesFn(c.ConfigureRoutes),
	)
	if err != nil {
		return
	}
	err = svc.Serve()
	return
}

func (s Service) Addr() *string {
	return env.StringOrNil(name + ".addr")
}

func (s Service) Kind() service.Kind {
	return http.Kind
}

func (s Service) CheckHealth(ctx context.Context) error {
	b := s.ClientBuilder
	if b == nil {
		b = service.DefaultHTTPClient
	}
	cli, err := b.NewHTTPClient(*s.Addr())
	if err != nil {
		return err
	}
	return http.CheckServiceHealth(ctx, cli, *s.Addr())
}
