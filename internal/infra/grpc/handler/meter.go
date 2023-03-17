package handler

import (
	"context"

	"enerBit-system/internal/app"
	"enerBit-system/internal/domain/dto"
	pb "enerBit-system/internal/infra/grpc/service"
	"github.com/andresxlp/gosuite/errs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcServer struct {
	pb.UnimplementedMeterServiceServer
	app app.Meter
}

func NewMeterServer(app app.Meter) pb.MeterServiceServer {
	return &grpcServer{
		app: app,
	}
}

func (handler *grpcServer) Create(ctx context.Context, req *pb.CreateMeterReq) (*pb.CreateMeterResp, error) {
	request := dto.Meter{
		Brand:  req.Meter.GetBrand(),
		Serial: req.Meter.GetSerial(),
		Lines:  int(req.Meter.GetLines()),
	}

	if err := request.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if err := handler.app.RegisterNewMeter(ctx, request); err != nil {
		if grpcErr, ok := err.(*errs.AppError); ok {
			return nil, grpcErr.NewGRPCError()
		}

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.CreateMeterResp{MeterId: "req.Meter."}, nil
}
