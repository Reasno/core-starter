package service

import (
	"context"

	pb "github.com/DoNewsCode/core-starter/app/proto"
)

type AppService struct {
	pb.UnimplementedAppServer
}

func NewAppService() pb.AppServer {
	return &AppService{}
}

func (s *AppService) CreateApp(ctx context.Context, req *pb.CreateAppRequest) (*pb.CreateAppReply, error) {
	return &pb.CreateAppReply{}, nil
}
func (s *AppService) UpdateApp(ctx context.Context, req *pb.UpdateAppRequest) (*pb.UpdateAppReply, error) {
	return &pb.UpdateAppReply{}, nil
}
func (s *AppService) DeleteApp(ctx context.Context, req *pb.DeleteAppRequest) (*pb.DeleteAppReply, error) {
	return &pb.DeleteAppReply{}, nil
}
func (s *AppService) GetApp(ctx context.Context, req *pb.GetAppRequest) (*pb.GetAppReply, error) {
	return &pb.GetAppReply{}, nil
}
func (s *AppService) ListApp(ctx context.Context, req *pb.ListAppRequest) (*pb.ListAppReply, error) {
	return &pb.ListAppReply{}, nil
}
