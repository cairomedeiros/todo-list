package handler

import "time"

type CreateToDoRequest struct {
	Id           uint      `json:"id" gorm:"primaryKey"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CreationDate time.Time `json:"creationDate"`
}
