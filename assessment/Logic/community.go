package Logic

import (
	"assessment/Dao/mysql"
	"assessment/Models"
)

// CreateCommunity 创建社区
func CreateCommunity(com *Models.Community) (err error) {
	return mysql.CreateCommunity(com)
}
