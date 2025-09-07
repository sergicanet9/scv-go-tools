package utils

import (
	"errors"

	"github.com/sergicanet9/scv-go-tools/v4/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ToGRPC maps errors to different gRPC errors with appropriate codes.
func ToGRPC(err error) error {
	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, wrappers.ValidationErr):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, wrappers.NonExistentErr):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, wrappers.UnauthorizedErr):
		return status.Error(codes.Unauthenticated, err.Error())
	case errors.Is(err, wrappers.UnauthenticatedErr):
		return status.Error(codes.PermissionDenied, err.Error())
	case errors.Is(err, wrappers.ServiceUnavailableErr):
		return status.Error(codes.Unavailable, err.Error())
	default:
		return status.Error(codes.Internal, err.Error())
	}
}
