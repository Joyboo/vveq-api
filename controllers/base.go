package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

//Response 结构体
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

//Response 结构体
type ErrResponse struct {
	Status int         `json:"status"`
	Msg    interface{} `json:"msg"`
}
