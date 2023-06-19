package rest

import (
	"github.com/cortezaproject/corteza/extra/server-discovery/pkg/auth"
	"github.com/cortezaproject/corteza/extra/server-discovery/searcher/rest/handlers"
)

func MountRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(auth.HttpTokenValidator("discovery"))
			handlers.NewSearch(Search()).MountRoutes(r)
		})
	}
}
