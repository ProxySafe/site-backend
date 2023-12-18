package repositories

type ErrDuplicateUser struct{}

func (e *ErrDuplicateUser) Error() string {
	return "user with such email/telephone/name already exists"
}
