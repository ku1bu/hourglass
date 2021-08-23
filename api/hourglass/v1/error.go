package v1

func (e *Errorgg) Error() string {
	return e.Code
}
