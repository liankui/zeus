package handlers

import (
	"context"
	"os/exec"

	"github.com/chaos-io/chaos/logs"

	// this service api
	pb "github.com/chaos-io/zeus/go/zeus/v1"
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

	soar := in.GetSoar()
	if soar == nil {
		return &pb.ExecSoarResponse{}, logs.NewErrorf("soar NOT be nil")
	}

	cmd := exec.Command("soar")
	if len(soar.Query) > 0 {
		cmd.Args = append(cmd.Args, "-query", soar.Query)
	}

	if len(soar.ReportType) > 0 {
		cmd.Args = append(cmd.Args, "-report", soar.ReportType)
	}

	logs.Infow("exec soar", "cmd", cmd.String())

	output, err := cmd.CombinedOutput()
	if err != nil {
		logs.Errorw("failed to exec soar", "cmd", cmd.String(), "error", err)
	}

	resp := &pb.ExecSoarResponse{
		// Level:
		Data: string(output),
	}
	return resp, nil
}
