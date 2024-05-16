package handler

import "time"

type CreateToDoRequest struct {
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CreationDate time.Time `json:"creationDate"`
}
