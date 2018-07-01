package golambdahelper

// GenerateState generates a new Shopify state value for Oauth
func GenerateState(uuid NewV4er) (string, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return uid.String(), nil
}
