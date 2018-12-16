package models

import (
	"github.com/astaxie/beego/orm"
	"regexp"
	"time"
	"unicode/utf8"
)

type User struct {
	Id            int
	Username      string `orm:"unique"`
	Nickname      string
	Password      string
	CheckPassword string `orm:"-"`
	Email         string
	Tel           string
	Avatar        string
	Instime       int64
}

func NewUser() *User {
	return &User{}
}

// 用户名是否存在
// @param username string
// @return int46
// @return error
func (u *User) GetUserByName(username string) (int64, error) {
	return orm.NewOrm().QueryTable("user").Filter("username", username).Count()
}

func (u *User) AddUser() (int64, error) {
	b, err := u.VerifyUserInfo()
	if err != nil || !b {
		return 0, err
	}
	u.Instime = time.Now().Unix()
	u.Password = Md5(u.Password)
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

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid int) {
	delete(UserList, uid)
}
*/
