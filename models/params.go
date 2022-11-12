package models

//header
type ParamHeader struct {
	Emp   string `header:"UUAP-EMPLOYEE-NUMBER" binding:"required,min=1" label:"请求header(工号)"`
	Email string `header:"UUAP-EMAIL" binding:"omitempty,min=1" label:"请求header(邮箱)"`
	Cname string `header:"UUAP-USERCNAME" binding:"omitempty,min=1" label:"请求header(用户中文名)"`
	Uname string `header:"UUAP-USERNAME" binding:"omitempty,min=1" label:"请求header(英文名)"`
}

type TestParams struct {
	Key string `form:"key" json:"key" binding:"omitempty,min=1" label:"关键字"`
}
