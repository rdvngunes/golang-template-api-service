package user

import (
	"github.com/google/uuid"
)

type UserCreateResponse struct {
	Id      uuid.UUID `json:"id"`
	Message string    `json:"message"`
}
