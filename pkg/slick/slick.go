package slick

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/a-h/templ"
)

type ErrorHandler func(error, *Context) error

type Context struct {
	response http.ResponseWriter
	request  *http.Request
	ctx      context.Context
}

func (c *Context) Render(component templ.Component) error {
	return component.Render(c.ctx, c.response)
}

type Handler func(c *Context) error

type Slick struct {
	ErrorHandler ErrorHandler
	router       *httprouter.Router
}

func New() *Slick {
	return &Slick{
		router:       httprouter.New(),
		ErrorHandler: defaultErrorHandler,
	}
}

func (s *Slick) Start(port string) error {
	return http.ListenAndServe(port, s.router)
}

func (s *Slick) Get(path string, handler Handler) {
	s.router.GET(path, s.makeHTTPRouterHandle(handler))
}

func (s *Slick) makeHTTPRouterHandle(handler Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		c := &Context{
			response: w,
			request:  r,
			ctx:      context.Background(),
		}
		if err := handler(c); err != nil {
			s.ErrorHandler(err, c)
		}
	}
}

func defaultErrorHandler(err error, c *Context) error {
	slog.Error("Error: %v", err)
	return err
}
