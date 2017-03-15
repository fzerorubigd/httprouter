package httprouter

import "context"

type ctxKey int

const (
	ctxKeyParam ctxKey = iota
	ctxPanicParam
)

// ParamsFromCtx return the param key from context
func ParamsFromCtx(ctx context.Context) Params {
	return ctx.Value(ctxKeyParam).(Params)
}

// PanicFromCtx return the panic key from context
func PanicFromCtx(ctx context.Context) interface{} {
	return ctx.Value(ctxPanicParam)
}

func withParam(parent context.Context, p Params) context.Context {
	return context.WithValue(parent, ctxKeyParam, p)
}

func withPanicCtx(parent context.Context, p interface{}) context.Context {
	return context.WithValue(parent, ctxPanicParam, p)
}
