package http

import (
	"context"
	"github.com/taranovegor/jurnalo/internal/domain"
	"github.com/taranovegor/jurnalo/internal/model/request"
	"github.com/taranovegor/jurnalo/internal/model/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type EntryHandler struct {
	Handler
	repository domain.EntryRepository
}

func NewEntryHandler(repository domain.EntryRepository) *EntryHandler {
	return &EntryHandler{
		repository: repository,
	}
}

func (h EntryHandler) List(rw http.ResponseWriter, req *http.Request) {
	paginator := request.Paginator{}
	parseQuery(req, &paginator)
	list, totalCount, err := h.repository.List(context.TODO(), paginator)
	if err != nil {
		errno(rw, http.StatusInternalServerError, err.Error())
	} else {
		resp(rw, http.StatusOK, response.NewPaginatedFromPaginator(paginator, totalCount, list))
	}
}

func (h EntryHandler) Get(rw http.ResponseWriter, req *http.Request) {
	id, err := primitive.ObjectIDFromHex(req.PathValue("id"))
	if err != nil {
		errno(rw, http.StatusBadRequest, err.Error())
	}

	entry, err := h.repository.Get(context.TODO(), id)
	if err != nil {
		errno(rw, http.StatusNotFound, err.Error())
	} else {
		resp(rw, http.StatusOK, entry)
	}
}
