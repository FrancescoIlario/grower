package valve

import (
	"context"

	"github.com/FrancescoIlario/grower/pkg/valvepb"
)

type valveServer struct {
	Cmder Commander
}

// NewGrpcServer ...
func NewGrpcServer(cmder Commander) valvepb.ValveServiceServer {
	return &valveServer{
		Cmder: cmder,
	}
}

func (v *valveServer) GetStatus(context.Context, *valvepb.GetStatusRequest) (*valvepb.GetStatusResponse, error) {
	st := v.Cmder.Status()
	statuspb := st.toStatusPB()
	return &valvepb.GetStatusResponse{Status: statuspb}, nil
}
func (v *valveServer) OpenValve(context.Context, *valvepb.OpenValveRequest) (*valvepb.OpenValveResponse, error) {
	v.Cmder.Open()
	return &valvepb.OpenValveResponse{}, nil
}

func (v *valveServer) CloseValve(context.Context, *valvepb.CloseValveRequest) (*valvepb.CloseValveResponse, error) {
	v.Cmder.Close()
	return &valvepb.CloseValveResponse{}, nil
}
