package Models

//UserForm 接收从前段传来的数据
type UserForm struct {
	UserName   string `json:"user_name" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required"`
}

//User user接收UserForm 存入mysql
type User struct {
	ID       uint64 `json:"id"`
	UserName string `json:"user_name"`
	UserID   int64  `json:"user_id"`
	Password string `json:"password"`
}
