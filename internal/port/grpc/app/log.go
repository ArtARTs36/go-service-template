package app

import (
	"context"

	"github.com/cappuccinotm/slogx/slogm"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

func injectRequestID() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		_ *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		ctx = slogm.ContextWithRequestID(ctx, uuid.New().String())

		return handler(ctx, req)
	}
}
