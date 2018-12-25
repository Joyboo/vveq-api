package models

import (
	"github.com/astaxie/beego/orm"
	"regexp"
	"time"
	"unicode/utf8"
)

type User struct {
	Id            int64
	Username      string `orm:"unique"`
	Nickname      string
	Password      string
	CheckPassword string `orm:"-"`
	Email         string
	Tel           string
	Avatar        string
	Instime       int64
	Status        int
}

func (d *User) TableName() string {
	return "user"
}

func NewUser() *User {
	return &User{}
}

// 用户名是否存在
// @param username string
// @return int46
// @return error
func (u *User) GetUserByName(username string) (int64, error) {
	return orm.NewOrm().QueryTable(u.TableName()).Filter("username", username).Count()
}

// 添加新用户
func (u *User) Add() (int64, error) {
	b, err := u.VerifyUserInfo()
	if err != nil || !b {
		return 0, err
	}
	num, err := u.GetUserByName(u.Username)
	if num > 0 || err != nil {
		return 0, err
	}
	u.Status = 1
	u.Instime = time.Now().Unix()
	u.Password = Md5(u.Password)
	// TODO 走redis队列
	return orm.NewOrm().Insert(u)
}

// 字段校验
func (u *User) VerifyUserInfo() (bool, error) {
	if u.Password != u.CheckPassword || utf8.RuneCountInString(u.Password) < 6 {
		return false, nil
	}
	isokUsername, err := regexp.MatchString(`^[a-zA-Z_\d]{4,20}$`, u.Username)
	if !isokUsername || err != nil {
		return false, err
	}
	isokEmail, err := regexp.MatchString(`^[0-9A-Za-z][\.-_0-9A-Za-z]*@[0-9A-Za-z]+(\.[A-Za-z]+)+$`, u.Email)
	if !isokEmail || err != nil {
		return false, err
	}
	return true, nil
}

func (u *User) GetUserById(id int64) (User, error) {
	var user User
	err := orm.NewOrm().QueryTable(u.TableName()).Filter("id", id).One(&user)
	return user, err
}

func (u *User) Login() (User, error) {
	var user User
	err := orm.NewOrm().QueryTable(u.TableName()).Filter("username", u.Username).Filter("password", Md5(u.Password)).Filter("status", 1).One(&user)
	return user, err
}

/*
func GetAllUsers() map[int]*User {
	return [0]User{}
}

func UpdateUser(uid int, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Address != "" {
			u.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func DeleteUser(uid int) {
	delete(UserList, uid)
}
*/
