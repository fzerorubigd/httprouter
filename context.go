package httprouter

import (
	"context"
	"net/http"
)

type ctxKey int

const (
	ctxKeyParam ctxKey = iota
	ctxPanicParam
)

// ParamsFromCtx return the param key from context
func ParamsFromCtx(r *http.Request) Params {
	return r.Context().Value(ctxKeyParam).(Params)
}

// PanicFromCtx return the panic key from context
func PanicFromCtx(r *http.Request) interface{} {
	return r.Context().Value(ctxPanicParam)
}

func withParam(parent *http.Request, p Params) *http.Request {
	ctx := context.WithValue(parent.Context(), ctxKeyParam, p)
	return parent.WithContext(ctx)
}

func withPanicCtx(parent *http.Request, p interface{}) *http.Request {
	ctx := context.WithValue(parent.Context(), ctxPanicParam, p)
	return parent.WithContext(ctx)
}
