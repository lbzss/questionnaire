package server

import (
	"fmt"
	"github.com/gorilla/handlers"
	"questionnaire/api/questionnaire/v1"
	"questionnaire/internal/conf"
	"questionnaire/internal/middleware"
	"questionnaire/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, q *service.QuestionnaireService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			middleware.Log(),
			recovery.Recovery(),
		),
		http.Filter(
			handlers.CORS(
				handlers.AllowedHeaders([]string{"Content-Type", "x-token"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}),
				handlers.AllowedOrigins([]string{"*"}))),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterQuestionnaireHTTPServer(srv, q)
	srv.WalkRoute(func(info http.RouteInfo) error {
		fmt.Printf("%-50s \t %s\n", info.Path, info.Method)
		return nil
	})
	return srv
}
