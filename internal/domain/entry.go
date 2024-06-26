package domain

import (
	"context"
	"github.com/mitchellh/mapstructure"
	"github.com/taranovegor/jurnalo/internal/model/request"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Entry map[string]interface{}

type EntryRepository interface {
	Store(context.Context, *Entry) error
	Get(context.Context, primitive.ObjectID) (*Entry, error)
	List(context.Context, request.Paginator) ([]Entry, int64, error)
}

func NewEntry(fields map[string]string, realtime time.Time) (*Entry, error) {
	var entry Entry
	err := mapstructure.Decode(fields, &entry)
	entry["_realtime"] = realtime.Format(time.RFC3339Nano)

	return &entry, err
}
