package join

import (
	"errors"
)

var (
	// ErrNotFound 該当オブジェクトが見つからなかった
	ErrNotFound = errors.New("not found")
)
