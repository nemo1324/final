package service

import (
	"context"
	finalv1 "final/pkg/proto/sync/final-boss/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) DeleteUser(ctx context.Context, req *finalv1.DeleteUserRequest) (*finalv1.DeleteUserResponse, error) {
	err := s.DB.DeleteUser(ctx, req.GetId())
	if err != nil {
		msg := "failed to delete user"
		s.logger.Error(msg+": %v", err)
		return nil, status.Error(codes.Internal, msg)
	}

	return &finalv1.DeleteUserResponse{}, nil
}
