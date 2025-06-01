package service

import (
	"context"
	"final/internal/repository/postgres/sqlc"
	finalv1 "final/pkg/proto/sync/final-boss/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) UpdateUser(ctx context.Context, req *finalv1.UpdateUserRequest) (*finalv1.UpdateUserResponse, error) {
	err := s.DB.UpdateUserName(ctx, sqlc.UpdateUserNameParams{
		ID:   int32(req.GetId()),
		Name: req.GetUsername(),
	})
	if err != nil {
		s.logger.Error("update user failed", "err", err)
		return nil, status.Error(codes.Internal, "update failed")
	}

	return &finalv1.UpdateUserResponse{}, nil
}
