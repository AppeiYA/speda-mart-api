package models

type Status string

const (
	StatusActive     Status = "active"
	StatusCheckedOut Status = "checked_out"
	StatusAbandoned  Status = "abandoned"
)

func (s Status) IsValid() bool {
	switch s {
	case StatusActive, StatusCheckedOut, StatusAbandoned:
		return true
	}
	return false
}

type Carts struct {
	Id        string `json:"id" db:"id"`
	UserId    string `json:"user_id" db:"user_id"`
	Status    Status `json:"status" db:"status"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
