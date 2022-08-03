package statuses

import (
	"encoding/json"
	"net/http"
	"strconv"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

// Handle request for `GET /v1/statuses/{id}`
func (h handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id_param := chi.URLParam(r, "id")
	var id int
	var err error
	if id_param == "" {
		httperror.BadRequest(w, errors.Errorf("id was not presence"))
	}
	id, err = strconv.Atoi(id_param)
	if err != nil {
		httperror.BadRequest(w, errors.Errorf("id is not a number"))
	}

	statusRepo := h.app.Dao.Status()
	var status *object.Status
	status, err = statusRepo.FindStatusByID(ctx, id)
	if err != nil {
		httperror.InternalServerError(w, err)
	}

	w.Header().Set("Content-Type", "applicatin/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		httperror.InternalServerError(w, err)
	}
}
