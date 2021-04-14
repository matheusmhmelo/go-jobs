package internal

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/matheusmhmelo/go-jobs/internal/handler"
	"github.com/matheusmhmelo/go-jobs/pkg/request"
	"net/http"
)

type Server struct {
	Handler http.Handler
}

func NewServer() http.Handler {
	r := mux.NewRouter()

	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	r.HandleFunc("/heartbeat", handler.Heartbeat).Methods(http.MethodGet)

	r.HandleFunc("/register", handler.Register).Methods(http.MethodPost)
	r.HandleFunc("/login", handler.Login).Methods(http.MethodPost)
	r.HandleFunc("/session", request.PreRequest(handler.ValidSession)).Methods(http.MethodGet)

	r.HandleFunc("/job", request.PreRequest(handler.CreateJob)).Methods(http.MethodPost)
	r.HandleFunc("/job/{id}", request.PreRequest(handler.GetJobInfo)).Methods(http.MethodGet)
	r.HandleFunc("/job/{id}", request.PreRequest(handler.UpdateJob)).Methods(http.MethodPut)
	r.HandleFunc("/job/{id}", request.PreRequest(handler.RemoveJob)).Methods(http.MethodDelete)

	return handlers.CORS(header, methods, origins)(r)
}