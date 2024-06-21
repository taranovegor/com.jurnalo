package container

import (
	"context"
	"github.com/coreos/go-systemd/v22/sdjournal"
	"github.com/sarulabs/di"
	"github.com/taranovegor/jurnalo/internal/config"
	"github.com/taranovegor/jurnalo/internal/handler/journal"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	JournalReader    = "journal_reader"
	MongoClient      = "mongodb_client"
	Database         = "mongodb_database"
	HandlerJournal   = "journal_handler"
	RepositoryEntry  = "repository_entry"
	HttpRouter       = "http_router"
	HttpEntryHandler = "http_entry_handler"
)

type ServiceContainer struct {
	container di.Container
}

func Init() (*ServiceContainer, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	return &ServiceContainer{
		container: build(builder),
	}, nil
}

func (sc ServiceContainer) Get(name string) interface{} {
	return sc.container.Get(name)
}

func build(builder *di.Builder) di.Container {
	builder.Add(di.Def{
		Name: JournalReader,
		Build: func(ctn di.Container) (interface{}, error) {
			return sdjournal.NewJournalReader(sdjournal.JournalReaderConfig{
				Since: time.Duration(-1) * time.Minute,
				Formatter: func(entry *sdjournal.JournalEntry) (string, error) {
					return ctn.Get(HandlerJournal).(*journal.Handler).Format(entry)
				},
			})
		},
	})

	builder.Add(di.Def{
		Name: MongoClient,
		Build: func(ctn di.Container) (interface{}, error) {
			apiOpts := options.ServerAPI(options.ServerAPIVersion1)
			opts := options.Client().ApplyURI(config.Get().Database.Uri).SetServerAPIOptions(apiOpts)

			ctx := context.Background()

			return mongo.Connect(ctx, opts)
		},
	})

	builder.Add(di.Def{
		Name: Database,
		Build: func(ctn di.Container) (interface{}, error) {
			return ctn.Get(MongoClient).(*mongo.Client).Database(config.Get().Database.Name), nil
		},
	})

	buildRepository(builder)
	buildHandler(builder)

	return builder.Build()
}
