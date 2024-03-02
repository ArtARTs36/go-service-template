package main

import (
	"log"
	"log/slog"

	"github.com/go-openapi/loads"

	"github.com/artarts36/go-service-template/internal/port/http/generated/restapi"
	apiOperations "github.com/artarts36/go-service-template/internal/port/http/generated/restapi/operations"

	"github.com/artarts36/go-service-template/internal/port/http"
	"github.com/artarts36/go-service-template/internal/port/http/app"
)

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	cfg, err := app.InitConfig("CARS_")
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	srv, err := app.New(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	api := apiOperations.NewCarsAPI(swaggerSpec)

	api.GetCarsIDHandler = apiOperations.GetCarsIDHandlerFunc(srv.GetCarsIDHandler)
	api.ServerShutdown = srv.OnShutdown
	server := restapi.NewServer(api, http.NewConfigurator())
	defer func() {
		shutdownErr := server.Shutdown()
		if shutdownErr != nil {
			slog.Error(shutdownErr.Error())
		}
	}()

	server.ConfigureAPI()

	server.Port = cfg.HTTP.Port
	if serveErr := server.Serve(); serveErr != nil {
		slog.Error(serveErr.Error())
	}
}
