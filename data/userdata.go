package data

type UserSign struct { //登陆用
	Username string `json:"username" gorm:"primary_key"` //用户名（学工号）
	Password string `json:"password"`                    //密码
}

type UserRegister struct { //注册用
	Identity string `json:"identity" `                   //权限，用户，管理员
	Username string `json:"username" gorm:"primary_key"` //用户名（学工号）
	Password string `json:"password" `                   //密码
	Email    string `json:"useremail" `                  //邮箱
}

func Adduser() {
	var user []UserRegister
	var single_user UserRegister
	user = append(user, single_user)
}
