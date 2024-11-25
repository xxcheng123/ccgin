package mysqlerr

import (
	"errors"

	mysql2 "github.com/go-sql-driver/mysql"
)

const (
	// ErrDuplicateEntryCode 命中唯一索引
	ErrDuplicateEntryCode = 1062
)

// ErrCode 根据mysql错误信息返回错误代码
func ErrCode(err error) int {
	var mysqlErr *mysql2.MySQLError
	ok := errors.As(err, &mysqlErr)

	if !ok {
		return 0
	}

	return int(mysqlErr.Number)
}

func IsDuplicatedKeyErr(err error) bool {
	return ErrCode(err) == ErrDuplicateEntryCode
}
