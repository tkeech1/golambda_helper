package golambda_helper

import (
	uuid "github.com/satori/go.uuid"
	"github.com/tkeech1/goshopify"
)

func (h *UuidHandler) Install(apiKey string, scope string, redirectUrl string, shopname string) (string, error) {
	uid, err := h.Uuid.NewV4()
	if err != nil {
		return "", err
	}
	state := uid.String()

	// TODO save state to database

	permissionUrl := goshopify.CreatePermissionUrl(apiKey, scope, redirectUrl, state, shopname)
	return permissionUrl, nil
}

func (t *UuidHandler) NewV4() (uuid.UUID, error) {
	return uuid.NewV4()
}
