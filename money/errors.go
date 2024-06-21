package money

// Error represents a money package error.
type Error string

// Error implements the error interface.
func (e Error) Error() string {
	return string(e)
}
