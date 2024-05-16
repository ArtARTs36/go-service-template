package app

import (
	"github.com/artarts36/go-service-template/internal/application/car"
	"github.com/go-openapi/runtime/middleware"

	"github.com/artarts36/go-service-template/internal/port/http/generated/models"
	apiOperations "github.com/artarts36/go-service-template/internal/port/http/generated/restapi/operations"
)

func (srv *Service) GetCarsIDHandler(params apiOperations.GetCarsIDParams) middleware.Responder {
	c, err := srv.container.Application.Operations.Car.Get.Get(
		params.HTTPRequest.Context(),
		&car.GetOperationParams{
			ID: params.ID,
		},
	)
	if err != nil {
		return apiOperations.NewGetCarsIDInternalServerError().WithPayload(&models.ErrorResponse{
			Message: err.Error(),
		})
	}

	return &apiOperations.GetCarsIDOK{
		Payload: &models.CarGetResponse{
			ID: c.ID,
		},
	}
}
