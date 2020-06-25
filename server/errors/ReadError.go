package errors

type ReadError struct {
	Message string
}

func (e *ReadError) Error() string {
	m := "Error reading"
	if e.Message!="" {
		m = e.Message
	}
	return m 
}