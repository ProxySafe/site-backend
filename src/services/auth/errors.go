package auth

type ErrTokenExpired struct {
}

func (e *ErrTokenExpired) Error() string {
	return "token has expired"
}

func IsErrTokenExpired(err error) bool {
	_, ok := err.(*ErrTokenExpired)
	return ok
}
