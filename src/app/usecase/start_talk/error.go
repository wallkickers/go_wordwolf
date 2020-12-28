package start_talk

import (
	"errors"
)

var (
	// ErrAlreadyStarted 既にスタートしている
	ErrAlreadyStarted = errors.New("is already Started")
)
