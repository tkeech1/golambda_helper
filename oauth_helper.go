package golambda_helper

import (
	uuid "github.com/satori/go.uuid"
)

func (h *UuidHandler) GenerateState() (string, error) {
	uid, err := h.Uuid.NewV4()
	if err != nil {
		return "", err
	}
	return uid.String(), nil

}

func (t *UuidHandler) NewV4() (uuid.UUID, error) {
	return uuid.NewV4()
}
