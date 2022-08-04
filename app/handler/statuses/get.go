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

	idParam := chi.URLParam(r, "id")
	var id int
	var err error
	if idParam == "" {
		httperror.BadRequest(w, errors.Errorf("id was not presence"))
	}
	id, err = strconv.Atoi(idParam)
	if err != nil {
		httperror.BadRequest(w, errors.Errorf("id is not a number"))
	}

	statusRepo := h.app.Dao.Status()
	var status *object.Status
	status, err = statusRepo.FindStatusByID(ctx, id)
	if err != nil {
		httperror.InternalServerError(w, err)
	}

	accountRepo := h.app.Dao.Account()
	var account *object.Account
	var accountId = status.AccountID
	account, err = accountRepo.FindAccountByID(ctx, accountId)
	if err != nil {
		httperror.InternalServerError(w, err)
	}
	status.Account = *account

	w.Header().Set("Content-Type", "applicatin/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		httperror.InternalServerError(w, err)
	}
}
