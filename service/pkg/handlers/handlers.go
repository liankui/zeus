package handlers

import (
	"context"

	"github.com/chaos-io/chaos/logs"

	// this service api
	pb "github.com/chaos-io/zeus/api/v1"
)

type zeusServer struct {
	pb.UnimplementedZeusServer
}

// NewService returns a naive, stateless implementation of Interface.
func NewService() pb.ZeusServer {
	return zeusServer{}
}

// ExecSoar implements Interface.
func (s zeusServer) ExecSoar(ctx context.Context, in *pb.ExecSoarRequest) (*pb.ExecSoarResponse, error) {
	logs.Infow("ExecSoar", "request", in)

	resp := &pb.ExecSoarResponse{
		// Level:
		// Data:
	}
	return resp, nil
}
