package error

type QueryError struct {
	Message string
}

func (e *QueryError) Error() string {
	m := "Query statement error"
	if e.Message!="" {
		m = e.Message
	}
	return m 
}