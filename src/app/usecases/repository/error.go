package repository

import (
	"errors"
)

var (
	// ErrNotFound 該当オブジェクトが見つからなかった
	ErrNotFound = errors.New("not found")

	// ErrIsExisted 該当オブジェクトは既に登録済み
	ErrIsExisted = errors.New("is Already Exist")
)
