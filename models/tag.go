package models

type Tag struct {
	Id      int64
	Name    string
	Status  int
	Insuid  int64
	Instime int64
}

func (t *Tag) TableName() string {
	return "theme"
}

func NewTag() *Tag {
	return &Tag{}
}

func (t *Tag) Add() (int64, error) {
	return 0, nil
}
