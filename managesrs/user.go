package managesrs

import "github.com/ajay-kumar-ks/gin-test/models"

type UserManager struct {
}

func NewUserManager() *UserManager{
	return &UserManager{}

}

func (userMgr *UserManager) Create(user *models.User) (*models.User, error){

	return nil, nil
}