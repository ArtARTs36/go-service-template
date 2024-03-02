package app

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/artarts36/go-service-template/internal/port/http/generated/models"
	apiOperations "github.com/artarts36/go-service-template/internal/port/http/generated/restapi/operations"
)

func (srv *Service) GetCarsIDHandler(_ apiOperations.GetCarsIDParams) middleware.Responder {
	return &apiOperations.GetCarsIDOK{
		Payload: &models.CarGetResponse{
			ID: 1,
		},
	}
}
