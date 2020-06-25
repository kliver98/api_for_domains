package errors

type NoPingError struct {
	Message string
}

func (e *NoPingError) Error() string {
	m := "Unreachable host"
	if e.Message!="" {
		m = e.Message
	}
	return m 
}