package app

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	"github.com/artarts36/go-service-template/internal/config"
	"github.com/artarts36/go-service-template/internal/port/grpc/handlers/cars"
	carsapi "github.com/artarts36/go-service-template/pkg/cars-grpc-api"
	log "github.com/sirupsen/logrus"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type App struct {
	gRPCServer *grpc.Server
	port       int
}

// NewApp creates new gRPC server app.
func NewApp(
	cfg *Config,
) *App {
	cont := config.InitContainer(&cfg.Config)

	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			// logging.StartCall, logging.FinishCall,
			logging.PayloadReceived, logging.PayloadSent,
		),
		// Add any other option (check functions starting with logging.With).
	}

	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) error {
			log.Error("Recovered from panic", slog.Any("panic", p))

			return status.Errorf(codes.Internal, "internal error")
		}),
	}

	gRPCServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(recoveryOpts...),
		logging.UnaryServerInterceptor(InterceptorLogger(), loggingOpts...),
	))

	carsapi.RegisterCarsServer(gRPCServer, cars.NewService(cont))

	return &App{
		gRPCServer: gRPCServer,
		port:       cfg.GRPC.Port,
	}
}

// InterceptorLogger adapts slog logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func InterceptorLogger() logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		log.WithContext(ctx).Info(msg)
	})
}

// Run runs gRPC server.
func (a *App) Run() error {
	const op = "grpcapp.Run"

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server started")

	if serveErr := a.gRPCServer.Serve(l); serveErr != nil {
		return fmt.Errorf("%s: %w", op, serveErr)
	}

	return nil
}

// Stop stops gRPC server.
func (a *App) Stop() {
	log.Info("[grpc][app] stopping gRPC server")

	a.gRPCServer.GracefulStop()
}
