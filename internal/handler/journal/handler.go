package journal

import (
	"context"
	"encoding/json"
	"github.com/coreos/go-systemd/v22/sdjournal"
	"github.com/taranovegor/jurnalo/internal/domain"
	"io"
	"time"
)

type Handler struct {
	io.Writer
	repository domain.EntryRepository
}

func NewHandler(
	repository domain.EntryRepository,
) *Handler {
	return &Handler{
		repository: repository,
	}
}

func (h Handler) Format(entry *sdjournal.JournalEntry) (string, error) {
	b, err := json.Marshal(entry)

	return string(b), err
}

func (h Handler) Write(b []byte) (int, error) {
	var entry sdjournal.JournalEntry
	err := json.Unmarshal(b, &entry)
	if err != nil {
		return 0, err
	}

	entr, err := domain.NewEntry(entry.Fields, time.Unix(0, int64(entry.RealtimeTimestamp)*int64(time.Microsecond)))
	if err != nil {
		return 0, err
	}

	err = h.repository.Store(context.TODO(), entr)
	if err != nil {
		return 0, err
	}

	return len(b), nil
}
