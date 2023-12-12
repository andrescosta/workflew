package main

import (
	"context"
	"log"

	"github.com/andrescosta/goico/pkg/service"
	pb "github.com/andrescosta/jobico/api/types"
	"github.com/andrescosta/jobico/srv/internal/queue"
)

const Name = "Queue"

func main() {
	svc, err := service.NewGrpcService(context.Background(), "queue",
		&pb.Queue_ServiceDesc,
		func(ctx context.Context) (any, error) {
			return queue.NewServer(ctx)
		})

	if err != nil {
		log.Panicf("error starting queue service because: %s", err)
	}
	if err = svc.Serve(); err != nil {
		log.Fatalf("error serving queue service  because: %s", err)
	}
}
