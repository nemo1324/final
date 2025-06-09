package service

import (
	"context"
	finalv1 "final/pkg/proto/sync/final-boss/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) DeleteUser(ctx context.Context, req *finalv1.DeleteUserRequest) (*finalv1.DeleteUserResponse, error) {
	s.logger.Debug("DeleteUser request received", "user_id", req.GetId())

	err := s.DB.DeleteUser(ctx, req.GetId())
	if err != nil {
		s.logger.Error("failed to delete user: %v", err)
		return nil, status.Error(codes.Internal, "failed to delete user")
	}

	return &finalv1.DeleteUserResponse{}, nil
}
