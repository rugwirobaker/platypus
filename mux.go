package platypus

import (
	"context"
	"strings"
)

type ctxKey string

const kparams ctxKey = "params"

var _ Handler = (*HandlerFunc)(nil)
var _ Handler = (*Mux)(nil)

// Handler ...
type Handler interface {
	Process(context.Context, *Command) (Result, error)
}

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as a Handler.
type HandlerFunc func(context.Context, *Command) (Result, error)

// Process calls fn(ctx, cmd)
func (fn HandlerFunc) Process(ctx context.Context, cmd *Command) (Result, error) {
	return fn(ctx, cmd)
}

// Mux ...
type Mux struct {
	tree     *node
	notFound Handler
}

//NewMux ...
func NewMux(prefix string, notFound Handler) *Mux {
	prefix = strings.TrimSuffix(prefix, "#")
	node := node{key: prefix, isParam: false}
	return &Mux{tree: &node, notFound: notFound}
}

// Process dispatches a command sequense to the handler whose
// pattern most closely matches the cmd pattern.
func (mux *Mux) Process(ctx context.Context, cmd *Command) (Result, error) {
	params := make(Params)

	cmd.Pattern = strings.TrimSuffix(cmd.Pattern, "#")

	node, _ := mux.tree.traverse(strings.Split(cmd.Pattern, "*")[1:], params)

	ctx = ContextWithParams(ctx, params)

	if node.action != nil {
		return node.action.Process(ctx, cmd)
	}
	return mux.notFound.Process(ctx, cmd)
}

// Handle ...
func (mux *Mux) Handle(pattern string, handler Handler) {
	pattern = strings.TrimSuffix(pattern, "#")

	if pattern[0] != '*' {
		panic("Path has to start with a *.")
	}
	if handler == nil {
		panic("mux: nil handler")
	}
	mux.tree.insertNode(pattern, handler)

}

// HandlerFunc registers the handler function for the given pattern.
func (mux *Mux) HandlerFunc(pattern string, handler func(context.Context, *Command) (Result, error)) {
	if pattern[0] != '*' {
		panic("Path has to start with a *.")
	}
	if handler == nil {
		panic("mux: nil handler")
	}
	mux.Handle(pattern, HandlerFunc(handler))
}

// NotFound returns an error indicating that the handler was not found for the given task.
func NotFound(ctx context.Context, cmd *Command) (Result, error) {
	return noResult("undefined"), nil
}

// NotFoundHandler returns a simple task handler that returns a ``not found`` error.
func NotFoundHandler() Handler { return HandlerFunc(NotFound) }
