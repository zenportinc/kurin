package http

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/zenportinc/kurin"
	httpAdapter "github.com/zenportinc/kurin/adapters/http"
	"github.com/zenportinc/kurin/example/engine"
	"go.uber.org/zap"
)

func NewAdapter(e engine.Engine, port int, logger *zap.Logger) kurin.Adapter {
	router := mux.NewRouter().StrictSlash(false)
	router.NewRoute().
		Name("List all users").
		Methods(http.MethodGet).
		Path("/users").
		Handler(listUsersHandler(e))
	router.NewRoute().
		Name("Create user").
		Methods(http.MethodPost).
		Path("/users").
		Handler(createUserHandler(e))
	router.NewRoute().
		Name("Get user").
		Methods(http.MethodGet).
		Path("/users/{id}").
		Handler(getUserHandler(e))
	router.NewRoute().
		Name("Delete user").
		Methods(http.MethodDelete).
		Path("/users/{id}").
		Handler(deleteUserHandler(e))

	handler := handlers.RecoveryHandler()(router)
	handler = handlers.CompressHandler(handler)
	handler = handlers.ContentTypeHandler(handler, "application/json")
	handler = handlers.CombinedLoggingHandler(os.Stdout, handler)

	return httpAdapter.NewHTTPAdapter(router, handler, port, "1.0.0", logger.Sugar())
}
