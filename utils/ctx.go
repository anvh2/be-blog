package utils

import (
	"context"
	"errors"

	"google.golang.org/grpc/metadata"
)

// GetUserIDFromCtx ...
func GetUserIDFromCtx(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("[Create] failed to get metadata from ctx")
	}

	val, ok := md["user_id"]
	if !ok || len(val) < 1 {
		return "", errors.New("failed to get userID")
	}

	return val[0], nil
}
