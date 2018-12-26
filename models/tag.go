package models

type Tag struct {
	Id      int64
	Name    string
	Status  int
	Insuid  int64
	Instime int64
}

func (this *Tag) TableName() string {
	return "tag"
}

func NewTag() *Tag {
	return &Tag{}
}

func (this *Tag) Add() (int64, error) {
	return 0, nil
}
