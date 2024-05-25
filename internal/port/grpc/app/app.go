package app

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"github.com/artarts36/go-service-template/internal/config"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
	container  *config.Container
}

// NewApp creates new gRPC server app.
func NewApp(
	cfg *Config,
	version string,
) (*App, error) {
	cont, err := config.InitContainer(&cfg.Config, version)
	if err != nil {
		return nil, err
	}

	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			logging.StartCall, logging.FinishCall,
			logging.PayloadReceived, logging.PayloadSent,
		),
		// Add any other option (check functions starting with logging.With).
	}

	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) error {
			slog.Error("Recovered from panic", slog.Any("panic", p))

			return status.Errorf(codes.Internal, "internal error")
		}),
	}

	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		injectRequestID(),
		recovery.UnaryServerInterceptor(recoveryOpts...),
		logging.UnaryServerInterceptor(InterceptorLogger(), loggingOpts...),
	))

	registerServices(gRPCServer, cont)

	if cfg.GRPC.UseReflection {
		reflection.Register(gRPCServer)
	}

	return &App{
		gRPCServer: gRPCServer,
		port:       cfg.GRPC.Port,
		container:  cont,
	}, nil
}

// InterceptorLogger adapts slog logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func InterceptorLogger() logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		switch lvl {
		case logging.LevelDebug:
			slog.DebugContext(ctx, msg, fields...)
		case logging.LevelInfo:
			slog.InfoContext(ctx, msg, fields...)
		case logging.LevelWarn:
			slog.WarnContext(ctx, msg, fields...)
		case logging.LevelError:
			slog.ErrorContext(ctx, msg, fields...)
		}
	})
}

// Run runs gRPC server.
func (a *App) Run() error {
	const op = "grpcapp.Run"

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	slog.Info("[app] grpc server started")

	if serveErr := a.gRPCServer.Serve(l); serveErr != nil {
		return fmt.Errorf("%s: %w", op, serveErr)
	}

	return nil
}

// Stop stops gRPC server.
func (a *App) Stop() {
	slog.Info("[app] stopping gRPC server")

	a.gRPCServer.GracefulStop()

	slog.Info("[app] closing db")

	if err := a.container.Infrastructure.DB.Close(); err != nil {
		slog.
			With(slog.Any("err", err)).
			Error("[app] failed to close db")
	}
}
