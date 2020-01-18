package service

import (
	"errors"
	"fastIM/app/model"
	"fastIM/app/util"
	"fmt"
	"math/rand"
	"time"
)

type UserService struct{}

//用户注册
func (s *UserService) UserRegister(mobile, plainPwd, nickname, avatar, sex string) (user model.User, err error) {
    registerUser := model.User{}
    _, err = model.DbEngine.Where("mobile=? ", mobile).Get(&registerUser)
    if err != nil {
    	return registerUser, err
	}
	//如果用户已经注册,返回错误信息
	if registerUser.Id > 0 {
		return registerUser, errors.New("该手机号已注册")
	}

	registerUser.Mobile = mobile
	registerUser.Avatar = avatar
	registerUser.Nickname = nickname
	registerUser.Sex = sex
	registerUser.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	registerUser.Passwd = util.MakePasswd(plainPwd, registerUser.Salt)
	registerUser.Createat = time.Now()
	//插入用户信息
	_, err = model.DbEngine.InsertOne(&registerUser)

	return registerUser,  err
}

//用户登录
func (s *UserService) Login(mobile, plainpwd string) (user model.User, err error) {
	//数据库操作
	loginUser := model.User{}
	model.DbEngine.Where("mobile = ?", mobile).Get(&loginUser)
	if loginUser.Id == 0 {
		return loginUser, errors.New("用户不存在")
	}
	//判断密码是否正确
	if !util.ValidatePasswd(plainpwd, loginUser.Salt, loginUser.Passwd) {
		return loginUser, errors.New("密码不正确")
	}
	//刷新用户登录的token值
	token := util.GenRandomStr(32)
	loginUser.Token = token
	model.DbEngine.ID(loginUser.Id).Cols("token").Update(&loginUser)

	//返回新用户信息
	return loginUser, nil
}

//查找某个用户
func (s *UserService) Find(userId int64) (user model.User) {
	findUser := model.User{}
	model.DbEngine.ID(userId).Get(&findUser)

	return findUser
}
