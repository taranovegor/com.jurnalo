package repository

import (
	"context"
	"github.com/taranovegor/jurnalo/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type entryRepository struct {
	domain.EntryRepository
	col *mongo.Collection
}

func NewEntryRepository(col *mongo.Collection) domain.EntryRepository {
	return entryRepository{
		col: col,
	}
}

func (r entryRepository) Store(ctx context.Context, link *domain.Entry) error {
	_, err := r.col.InsertOne(ctx, link)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	return err
}

func (r entryRepository) Get(ctx context.Context, id primitive.ObjectID) (*domain.Entry, error) {
	var entry domain.Entry
	err := r.col.FindOne(ctx, bson.D{{"_id", id}}).Decode(&entry)

	return &entry, err
}

func (r entryRepository) List(ctx context.Context, paginator domain.Paginator) ([]domain.Entry, int64, error) {
	var entries []domain.Entry

	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"_id", -1}}).SetSkip(paginator.Offset).SetLimit(paginator.Limit)
	cur, err := r.col.Find(ctx, filter, opts)
	if err != nil {
		return entries, 0, err
	}

	err = cur.All(ctx, &entries)
	if err != nil {
		return entries, 0, err
	}

	totalCount, err := r.col.CountDocuments(ctx, filter)
	if err != nil {
		return entries, 0, err
	}

	return entries, totalCount, nil
}
