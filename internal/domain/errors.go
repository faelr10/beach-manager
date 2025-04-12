package domain

import "errors"

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrUserAlreadyExists   = errors.New("user already exists")
	ErrInvalidInput        = errors.New("invalid input")
	ErrInternalServerError = errors.New("internal server error")

	ErrAgendaNotFound      = errors.New("agenda not found")
	ErrAgendaAlreadyExists = errors.New("agenda already exists")
	ErrAgendaConflict      = errors.New("agenda conflict")
	ErrAgendaInvalidDate   = errors.New("invalid date format")
	ErrAgendaInvalidTime   = errors.New("invalid time format")

	ErrAuthInvalidCredentials = errors.New("invalid credentials")
)
