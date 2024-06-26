package http

import (
	"github.com/taranovegor/jurnalo/internal/domain"
	"net/http"
)

type AuthHandler struct {
	Handler
	repository domain.UserRepository
}

func NewAuthHandler(repository domain.UserRepository) *AuthHandler {
	return &AuthHandler{
		repository: repository,
	}
}

func (h AuthHandler) Login(rw http.ResponseWriter, req *http.Request) {
	//id, err := primitive.ObjectIDFromHex(req.PathValue("id"))
	//if err != nil {
	//	errno(rw, http.StatusBadRequest, err.Error())
	//}
	//
	//entry, err := h.repository.Get(context.TODO(), id)
	//if err != nil {
	//	errno(rw, http.StatusNotFound, err.Error())
	//} else {
	//	resp(rw, http.StatusOK, entry)
	//}
}
