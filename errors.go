package berrors

import "fmt"

// Errorf wrapps error with message using fmt.Errorf if err is not nil
func Errorf(err error, msg string) error {
	if err != nil {
		return fmt.Errorf(msg, err)
	}
	return nil
}
