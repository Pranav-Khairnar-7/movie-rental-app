package model

import uuid "github.com/satori/go.uuid"

type UserMovies struct {
	Base
	UserID  uuid.UUID
	MovieID uuid.UUID
}
