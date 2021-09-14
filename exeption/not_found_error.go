package exeption

type NotFoundError struct {
	Err string
}

func NewNotFoundError(err string) *NotFoundError {
	return &NotFoundError{Err: err}
}

func (nt *NotFoundError) Error() string {
	return nt.Err
}
