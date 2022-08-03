package object

type (
	StatusID = int64

	Status struct {
		ID StatusID `json:"-"`

		AccountID AccountID `json:"account_id,omitempty" db:"account_id"`

		Status string `json:"status" db:"content"`

		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}
)
