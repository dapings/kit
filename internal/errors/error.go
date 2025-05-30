package errors

import "fmt"

func NewIndexOutOfRange(length, index int) error {
	return fmt.Errorf("index out of range: length=%d, index=%d", length, index)
}
