// This file is safe to edit. Once it exists it will not be overwritten

package http

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"github.com/artarts36/go-service-template/internal/port/http/generated/restapi/operations"
	"github.com/artarts36/go-service-template/internal/port/http/middlewares"
)

type Configurator struct {
}

func NewConfigurator() *Configurator {
	return &Configurator{}
}

func (c *Configurator) ConfigureFlags(_ *operations.CarsAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func (c *Configurator) ConfigureAPI(api *operations.CarsAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.PreServerShutdown = func() {}
	api.ServerShutdown = func() {}

	return c.SetupGlobalMiddleware(api.Serve(c.SetupMiddlewares))
}

// ConfigureTLS The TLS configuration before HTTPS server starts.
func (c *Configurator) ConfigureTLS(_ *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// ConfigureServer As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func (c *Configurator) ConfigureServer(_ *http.Server, _, _ string) {
}

// SetupMiddlewares The middleware configuration is for the handler executors.
// These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func (c *Configurator) SetupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// SetupGlobalMiddleware The middleware configuration happens before anything.
// This middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func (c *Configurator) SetupGlobalMiddleware(handler http.Handler) http.Handler {
	return middlewares.CORS(middlewares.NewLog(handler))
}
