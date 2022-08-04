package timelines

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/handler/httperror"
)

// Handle request for `GET /v1/timelines/public`
func (h handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// var err error

	// max_ids := r.FormValue("max_id")
	// var max_id int
	// if max_ids == "" {
	// 	max_id = 0
	// } else {
	// 	max_id, err = strconv.Atoi(max_ids)
	// 	if err != nil {
	// 		httperror.BadRequest(w, errors.Errorf("max_id is not a number"))
	// 	}
	// }

	// since_ids := r.FormValue("since_id")
	// var since_id int
	// if since_ids == "" {
	// 	since_id = 0
	// } else {
	// 	since_id, err = strconv.Atoi(since_ids)
	// 	if err != nil {
	// 		httperror.BadRequest(w, errors.Errorf("since_id is not a number"))
	// 	}
	// }

	// if max_id < since_id {
	// 	httperror.BadRequest(w, errors.Errorf("since_id is bigger than max_id"))
	// }

	// limit_s := r.FormValue("limit")
	// var limit int
	// if limit_s == "" {
	// 	limit = 40
	// } else {
	// 	limit, err = strconv.Atoi(since_ids)
	// 	if err != nil {
	// 		httperror.BadRequest(w, errors.Errorf("since_id is not a number"))
	// 	}
	// }
	// if limit > 80 {
	// 	limit = 80
	// }

	statusRepo := h.app.Dao.Status()
	timelines, err := statusRepo.FindAllStatus(ctx)
	if err != nil {
		httperror.InternalServerError(w, err)
	}

	accountRepo := h.app.Dao.Account()
	for _, s := range timelines {
		account_id := s.AccountID
		account, err := accountRepo.FindAccountByID(ctx, account_id)
		if err != nil {
			httperror.InternalServerError(w, err)
			return
		}
		s.Account = *account
	}

	w.Header().Set("Content-Type", "applicatin/json")
	if err := json.NewEncoder(w).Encode(timelines); err != nil {
		httperror.InternalServerError(w, err)
	}
}
