package platypus

import "context"

// Params named params
type Params map[string]interface{}

// Add paramater
func (p Params) Add(key string, val interface{}) {
	p[key] = val
}

// GetInt returns an integer value from params
func (p Params) GetInt(key string) int {
	return p[key].(int)
}

// GetString returns an string value from params
func (p Params) GetString(key string) string {
	return p[key].(string)
}

// GetBool returns an bool value from params
func (p Params) GetBool(key string) bool {
	return p[key].(bool)
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
