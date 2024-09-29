package handlers

import "github.com/ajay-kumar-ks/gin-test/managesrs"

type UserHandler struct {
	groupName string
	userManager *managesrs.UserManager
}

func NewUserHanlderFrom() *UserHandler{
	return &UserHandler{
		"user",
	    userManager,
	}
	
}

func (userHandler *UserHandler) RegisterUserapis(){
}
func (userHandler *UserHandler) Create(){
}