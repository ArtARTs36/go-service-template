package main

import (
	"github.com/go-openapi/loads"
	log "github.com/sirupsen/logrus"

	"github.com/artarts36/go-service-template/internal/port/http/generated/restapi"
	apiOperations "github.com/artarts36/go-service-template/internal/port/http/generated/restapi/operations"

	"github.com/artarts36/go-service-template/internal/port/http/app"
)

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	cfg, err := app.InitConfig("CARS_")
	if err != nil {
		log.Fatalln(err)
	}

	srv := app.New(cfg)
	api := apiOperations.NewCarsAPI(swaggerSpec)

	api.GetCarsIDHandler = apiOperations.GetCarsIDHandlerFunc(srv.GetCarsIDHandler)
	api.ServerShutdown = srv.OnShutdown
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.ConfigureAPI()

	api.Logger = log.Debugf

	server.Port = cfg.HTTP.Port
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
