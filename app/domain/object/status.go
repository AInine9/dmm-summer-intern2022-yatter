package object

type (
	StatusID = int64

	Status struct {
		ID StatusID `json:"id,omitempty"`

		AccountID AccountID `json:"-" db:"account_id"`

		Account Account `json:"account"`

		Status string `json:"content" db:"content"`

		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}
)
