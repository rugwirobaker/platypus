package platypus

import (
	"context"
	"fmt"
)

// Params named params
type Params map[string]interface{}

type keyError struct {
	key string
}

func (e *keyError) Error() string {
	return fmt.Sprintf("key %q does not exist", e.key)
}

// Add paramater
func (p Params) Add(key string, val interface{}) {
	p[key] = val
}

// GetInt returns an integer value from params
func (p Params) GetInt(key string) (int, error) {
	if val, ok := p[key]; ok {
		return val.(int), nil
	}
	return 0, &keyError{key}
}

// GetString returns an string value from params
func (p Params) GetString(key string) (string, error) {
	if val, ok := p[key]; ok {
		return val.(string), nil
	}
	return "", &keyError{key}
}

// GetBool returns an bool value from params
func (p Params) GetBool(key string) (bool, error) {
	if val, ok := p[key]; ok {
		return val.(bool), nil
	}
	return false, &keyError{key}
}

// ContextWithParams ...
func ContextWithParams(ctx context.Context, val Params) context.Context {
	if val == nil {
		return ctx
	}
	return context.WithValue(ctx, kparams, val)
}

// ParamsFromContext extracts params from context
func ParamsFromContext(ctx context.Context) Params {
	e, ok := ctx.Value(kparams).(Params)
	if !ok || e == nil {
		return make(Params)
	}
	return e
}
