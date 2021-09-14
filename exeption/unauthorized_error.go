package exeption

type UnAuthorizedError struct {
	Err string
}

func NewUnAuthorizedError(err string) *UnAuthorizedError {
	return &UnAuthorizedError{Err: err}
}

func (nt *UnAuthorizedError) Error() string {
	return nt.Err
}
