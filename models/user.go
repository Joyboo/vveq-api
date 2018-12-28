package models

import (
	"github.com/astaxie/beego/orm"
	"regexp"
	"time"
	"unicode/utf8"
)

type User struct {
	Id            int64  `json:"id"`
	Username      string `orm:"unique" json:"username"`
	Nickname      string `json:"nickname"`
	Password      string
	CheckPassword string    `orm:"-"`
	Email         string    `json:"email"`
	Tel           string    `json:"tel"`
	Avatar        string    `json:"avatar"`
	Status        int       `json:"status"`
	Instime       time.Time `orm:"column(instime);auto_now_add;type(datetime)" json:"instime"`
	Updtime       time.Time `orm:"column(updtime);auto_now;type(datetime)" json:updatime`
}

func NewUser() *User {
	return &User{}
}

func (this *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(User)).Filter("status", 1)
}

// 用户名是否存在
// @param username string
// @return int46
// @return error
func (this *User) UsernameExist(username string) bool {
	return this.Query().Filter("username", username).Exist()
}

// 添加新用户
func (this *User) Add() (int64, error) {
	b, err := this.VerifyUserInfo()
	if err != nil || !b {
		return 0, err
	}
	exist := this.UsernameExist(this.Username)
	if exist || err != nil {
		return 0, err
	}
	this.Status = 1
	this.Password = Md5(this.Password)
	// TODO 走redis队列
	return orm.NewOrm().Insert(this)
}

// 字段校验
func (this *User) VerifyUserInfo() (bool, error) {
	if this.Password != this.CheckPassword || utf8.RuneCountInString(this.Password) < 6 {
		return false, nil
	}
	isokUsername, err := regexp.MatchString(`^[a-zA-Z_\d]{4,20}$`, this.Username)
	if !isokUsername || err != nil {
		return false, err
	}
	isokEmail, err := regexp.MatchString(`^[0-9A-Za-z][\.-_0-9A-Za-z]*@[0-9A-Za-z]+(\.[A-Za-z]+)+$`, this.Email)
	if !isokEmail || err != nil {
		return false, err
	}
	return true, nil
}

func (this *User) GetUserById(id int64) (User, error) {
	var user User
	err := this.Query().Filter("id", id).One(&user)
	return user, err
}

func (this *User) Login() (User, error) {
	var user User
	err := this.Query().Filter("username", this.Username).Filter("password", Md5(this.Password)).One(&user)
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
