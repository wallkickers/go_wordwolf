package repository

import (
	"errors"
)

// ErrNotFound 該当オブジェクトが見つからなかった
var ErrNotFound = errors.New("not found")

// ErrIsExisted 該当オブジェクトは既に登録済み
var ErrIsExisted = errors.New("is Already Exist")
