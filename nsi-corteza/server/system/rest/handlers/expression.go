package handlers

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
//

import (
    "context"
    "github.com/cortezaproject/corteza/server/pkg/api"
    "github.com/cortezaproject/corteza/server/system/rest/request"
    "net/http"
)

type (
	// Internal API interface
	ExpressionAPI interface {
		Evaluate(context.Context, *request.ExpressionEvaluate) (interface{}, error)
	}

	// HTTP API interface
	Expression struct {
		Evaluate func(http.ResponseWriter, *http.Request)
	}
)

func NewExpression(h ExpressionAPI) *Expression {
	return &Expression{
		Evaluate: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewExpressionEvaluate()
			if err := params.Fill(r); err != nil {
				api.Send(w, r, err)
				return
			}

			value, err := h.Evaluate(r.Context(), params)
			if err != nil {
				api.Send(w, r, err)
				return
			}

			api.Send(w, r, value)
		},
	}
}

func (h Expression) MountRoutes(r chi.Router, middlewares ...func(http.Handler) http.Handler) {
	r.Group(func(r chi.Router) {
		r.Use(middlewares...)
		r.Post("/expressions/evaluate", h.Evaluate)
	})
}
