package Logic

import (
	"assessment/Dao/mysql"
	"assessment/Models"
	"assessment/pkg/snowflake"
)

//SignUp 注册
func SignUp(userform *Models.UserForm) (err error) {
	userID := snowflake.GenID()
	user := &Models.User{
		UserID:   userID,
		UserName: userform.UserName,
		Password: userform.Password,
	}
	err = mysql.SignUp(user)
	return
}

//Login 登陆
func Login(user *Models.User) (err error) {
	err = mysql.Login(user)
	return
}
