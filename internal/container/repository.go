package container

import (
	"github.com/sarulabs/di"
	"github.com/taranovegor/jurnalo/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func buildRepository(builder *di.Builder) {
	builder.Add(di.Def{
		Name: RepositoryEntry,
		Build: func(ctn di.Container) (interface{}, error) {
			return repository.NewEntryRepository(
				ctn.Get(Database).(*mongo.Database).Collection("entry"),
			), nil
		},
	})
}
