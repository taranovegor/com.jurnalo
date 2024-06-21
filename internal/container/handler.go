package container

import (
	"github.com/sarulabs/di"
	"github.com/taranovegor/jurnalo/internal/domain"
	"github.com/taranovegor/jurnalo/internal/handler/http"
	"github.com/taranovegor/jurnalo/internal/handler/journal"
	mux "net/http"
)

func buildHandler(builder *di.Builder) {
	builder.Add(di.Def{
		Name: HandlerJournal,
		Build: func(ctn di.Container) (interface{}, error) {
			return journal.NewHandler(
				ctn.Get(RepositoryEntry).(domain.EntryRepository),
			), nil
		},
	})

	builder.Add(di.Def{
		Name: HttpEntryHandler,
		Build: func(ctn di.Container) (interface{}, error) {
			return http.NewEntryHandler(
				ctn.Get(RepositoryEntry).(domain.EntryRepository),
			), nil
		},
	})

	builder.Add(di.Def{
		Name: HttpRouter,
		Build: func(ctn di.Container) (interface{}, error) {
			router := mux.NewServeMux()

			entryHandler := ctn.Get(HttpEntryHandler).(*http.EntryHandler)
			router.HandleFunc("GET /api/entry", entryHandler.List)
			router.HandleFunc("GET /api/entry/{id}", entryHandler.Get)

			return router, nil
		},
	})
}
