package models

type Theme struct {
	Id            int64
	Cid           int
	Title         string
	Content       string
	Uid           int64
	Tagid         int
	Sort          int
	Click         int64
	Like          int64
	Instime       int64
	Lastreplytime int64
}

func (t *Theme) TableName() string {
	return "theme"
}

func NewTheme() *Theme {
	return &Theme{}
}

func (t *Theme) Add() (int64, error) {
	return 0, nil
}
