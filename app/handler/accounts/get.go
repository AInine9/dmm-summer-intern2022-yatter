package accounts

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

//Handle request for `GET /v1/accounts/{username}`
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := chi.URLParam(r, "username")
	if username == "" {
		httperror.BadRequest(w, errors.Errorf("username was not presence"))
	}

	accountRepo := h.app.Dao.Account()
	var account *object.Account
	var err error
	account, err = accountRepo.FindByUsername(ctx, username)
	if err != nil {
		httperror.InternalServerError(w, err)
	}

	w.Header().Set("Content-Type", "applicatin/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		httperror.InternalServerError(w, err)
	}
}
