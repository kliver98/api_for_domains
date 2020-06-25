package errors

type ExecError struct {
	Message string
}

func (e *ExecError) Error() string {
	m := "Execution error"
	if e.Message!="" {
		m = e.Message
	}
	return m 
}