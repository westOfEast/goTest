package common

type Error struct {
	Msg string
	Err error
}

func (e *Error) Error() string {
	return e.Msg + e.Err.Error()
}
