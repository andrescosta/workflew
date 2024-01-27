package remote

import (
	"context"

	"github.com/andrescosta/goico/pkg/env"
	"github.com/andrescosta/goico/pkg/service"
	pb "github.com/andrescosta/jobico/internal/api/types"
	"github.com/rs/zerolog"
	rpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RecorderClient struct {
	serverAddr string
	conn       *rpc.ClientConn
	client     pb.RecorderClient
}

func NewRecorderClient(ctx context.Context, d service.GrpcDialer) (*RecorderClient, error) {
	addr := env.String("recorder.host")
	conn, err := d.Dial(ctx, addr)
	if err != nil {
		return nil, err
	}
	client := pb.NewRecorderClient(conn)
	return &RecorderClient{
		serverAddr: addr,
		conn:       conn,
		client:     client,
	}, nil
}

func (c *RecorderClient) Close() error {
	return c.conn.Close()
}

func (c *RecorderClient) StreamJobExecutions(ctx context.Context, lines int32, resChan chan<- string) error {
	logger := zerolog.Ctx(ctx)
	rj, err := c.client.GetJobExecutionsStr(ctx, &pb.GetJobExecutionsRequest{
		Lines: &lines,
	})
	if err != nil {
		return err
	}
	for {
		select {
		case <-rj.Context().Done():
			if err := rj.CloseSend(); err != nil {
				logger.Warn().AnErr("err", err).Msg("Recorder Client: error closing client stream")
			}
			return nil
		default:
			ress, err := rj.Recv()
			if err != nil {
				if status.Code(err) != codes.Canceled {
					logger.Warn().Msgf("error getting message %s", err)
				}
			} else {
				for _, r := range ress.Result {
					resChan <- r
				}
			}
		}
	}
}

func (c *RecorderClient) AddJobExecution(ctx context.Context, ex *pb.JobExecution) error {
	if _, err := c.client.AddJobExecution(ctx, &pb.AddJobExecutionRequest{Execution: ex}); err != nil {
		return err
	}
	return nil
}

func (c *RecorderClient) GetJobExecutions(ctx context.Context, tenant string, lines int32) ([]string, error) {
	res, err := c.client.GetJobExecutions(ctx, &pb.GetJobExecutionsRequest{
		Tenant: &tenant,
		Lines:  &lines,
	})
	if err != nil {
		return nil, err
	}
	return res.Result, nil
}
