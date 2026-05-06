package handler

import "github.com/go-core-fx/telegofx"

// Handler registers bot routes in the shared telegofx router.
type Handler interface {
	Register(router *telegofx.Router)
}
