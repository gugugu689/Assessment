package mysql

import (
	"assessment/Models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"

	"go.uber.org/zap"
)

const secret = "保命护身"

//encryptPassword 密码加密
func encryptPassword(data []byte) (result string) {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum(data))
}

//SignUp 注册
func SignUp(user *Models.User) (err error) {
	var count int
	db.Model(&Models.User{}).Where("user_id=?", user.UserID).Count(&count)
	if err != nil && err != sql.ErrNoRows {
		zap.L().Error("查询错误", zap.Error(err))
		return err
	}
	if count > 0 {
		return ErrorUserExit
	}
	//加密密码
	password := encryptPassword([]byte(user.Password))
	u := &Models.User{
		UserID:   user.UserID,
		UserName: user.UserName,
		Password: password,
	}
	db.Create(u)
	return
}

//Login 登陆
func Login(user *Models.User) (err error) {
	// 记录原始密码
	originPassword := user.Password
	//定义接收查询结果的结构体变量
	User := &Models.User{}
	//查询错误
	if err = db.Where("user_name=?", user.UserName).Find(User).Error; err != nil && err != sql.ErrNoRows {
		zap.L().Error("查询错误", zap.Error(err))
		return
	}
	//用户不存在
	if err == sql.ErrNoRows {
		return ErrorUserNotExit
	}
	//生成加密密码与查询到的密码比较
	if User.Password != encryptPassword([]byte(originPassword)) {
		return ErrorPasswordWrong
	}
	return
}

//GetUserByID 通过userid获取user
func GetUserByID(id int64) (user *Models.User, err error) {
	user = new(Models.User)
	if err = db.Where("user_id=?", id).Find(user).Error; err != nil {
		zap.L().Error("GetUserByID failed", zap.Error(err))
	}
	return
}
