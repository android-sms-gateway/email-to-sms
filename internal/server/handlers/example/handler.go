package example

import (
	"github.com/capcom6/go-project-template/internal/example"
	"github.com/go-core-fx/fiberfx/handler"
	"github.com/go-core-fx/fiberfx/validation"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Handler exposes example endpoints over HTTP.
type Handler struct {
	handler.Base

	exampleSvc *example.Service
}

// New creates a new example HTTP handler.
func New(service *example.Service, validator *validator.Validate) handler.Handler {
	return &Handler{
		Base: handler.Base{
			Validator: validator,
		},
		exampleSvc: service,
	}
}

// Register registers all example HTTP routes.
func (h *Handler) Register(r fiber.Router) {
	g := r.Group("/example")

	g.Get("/", h.get)
	g.Post("/", validation.DecorateWithBodyEx(h.Validator, h.post))
}

//	@Summary		Get example
//	@Description	Returns a greeting message and example value
//	@Tags			example
//	@Produce		json
//	@Success		200	{object}	Response
//	@Failure		500	{object}	fiberfx.ErrorResponse
//	@Router			/example [get]
//
// Get example.
func (h *Handler) get(c *fiber.Ctx) error {
	return c.JSON(Response{
		Message: "example handler is working",
		Value:   h.exampleSvc.Example(),
	})
}

//	@Summary		Create example
//	@Description	Creates an example using the provided request payload
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Param			request	body		Request	true	"Request payload"
//	@Success		200		{object}	Response
//	@Failure		400		{object}	fiberfx.ErrorResponse
//	@Failure		500		{object}	fiberfx.ErrorResponse
//	@Router			/example [post]
//
// Create example.
func (h *Handler) post(c *fiber.Ctx, req *Request) error {
	return c.JSON(Response{
		Message: "example handler is working",
		Value:   req.Value + " / " + h.exampleSvc.Example(),
	})
}
