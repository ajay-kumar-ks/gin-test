package managesrs

import (
	"errors"

	"github.com/ajay-kumar-ks/gin-test/common"
	"github.com/ajay-kumar-ks/gin-test/database"
	"github.com/ajay-kumar-ks/gin-test/models"
)

type UserManager struct {
}

func NewUserManager() *UserManager{
	return &UserManager{}

}

func (userMgr *UserManager) Create(userData *common.UserCreationInput) (*models.User, error){

	newUser := &models.User{FullName: userData.FullName, Email: userData.Email}
	database.DB.Create(newUser)
	if newUser.ID == 0 {
		return nil, errors.New("failed to create user")
	}
	return newUser, nil
}