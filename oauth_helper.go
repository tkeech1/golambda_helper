package golambda_helper

func GenerateState(uuid NewV4er) (string, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return uid.String(), nil
}
