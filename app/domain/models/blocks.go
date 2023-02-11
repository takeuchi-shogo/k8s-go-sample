package models

type Blocks struct {
	ID       int   `json:"id"`
	Blocking int   `json:"blocking"`
	Blocked  int   `json:"blocked"`
	CreateAt int64 `json:"created_at"`
}
