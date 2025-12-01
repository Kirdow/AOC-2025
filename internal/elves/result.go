package elves

import "fmt"

func Ok[T any](value T) (string, error) {
	return fmt.Sprintf("%v", value), nil
}

func Err(err error) (string, error) {
	return "", err
}
