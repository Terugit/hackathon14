package model

import "errors"

var (
	EmptySearchName       = errors.New("search name's empty")
	IncorrectRegisterName = errors.New("register name is not correct")
	FailIdGenerate        = errors.New("fail to build id")
	EmptyId               = errors.New("id empty")
	EmptyName             = errors.New("name empty")
	EmptyMessage          = errors.New("message empty")
	IncorrectPoint        = errors.New("point is not correct")
)

//IncorrectRegisterAge  = errors.New("register age is not correct")
