package models

type Cate struct {
	Id      int64
	Pid     int
	Name    string
	Sort    int
	Route   string
	Status  int
	Insuid  int64
	Instime int64
	Upduid  int64
	Updtime int64
}

func (t *Cate) TableName() string {
	return "theme"
}

func NewCate() *Cate {
	return &Cate{}
}

func (t *Cate) Add() (int64, error) {
	return 0, nil
}
