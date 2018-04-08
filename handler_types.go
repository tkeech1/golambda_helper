package golambda_helper

import (
	uuid "github.com/satori/go.uuid"
)

type UUIDInterface interface {
	NewV4() (uuid.UUID, error)
}

type UuidHandler struct {
	Uuid UUIDInterface
}
