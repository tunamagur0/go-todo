package models

import "time"

type TodoStatus int

const (
	StatusUnknown TodoStatus = iota
	StatusNew
	StatusWIP
	StatusDone
	StatusPending
)

func (ts TodoStatus) String() string {
	switch ts {
	case StatusNew:
		return "new"
	case StatusWIP:
		return "wip"
	case StatusDone:
		return "done"
	case StatusPending:
		return "pending"
	default:
		return "unknown"
	}
}

func NewTodoStatus(status int) TodoStatus {
	if status == 1 {
		return StatusNew
	}
	if status == 2 {
		return StatusWIP
	}
	if status == 3 {
		return StatusDone
	}
	if status == 4 {
		return StatusWIP
	}
	return StatusUnknown
}

type Todo struct {
	ID         string     `json:"id" gorm:"primary_key;type:text"`
	Content    string     `json:"content" gorm:"type:text"`
	Status     TodoStatus `json:"status"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"-"`
	FinishedAt *time.Time `json:"finished_at"`
}
