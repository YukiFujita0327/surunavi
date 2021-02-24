package gateways

import (
	"gorm.io/gorm"
	"surunavi/go/pkg/adapter/gateways/db"
	"surunavi/go/pkg/domain"
)

type (
	LoginRepository struct {
		Conn *gorm.DB
	}
)

func (repository *LoginRepository) GetUserInfo(userId string) (domain.UserInfo, error) {
	userInfo := &db.UserInfo{}
	if err := repository.Conn.Where("Id = ?", userId).First(&userInfo).Error
	err != nil {
		return domain.UserInfo{}, err
	}
	return domain.UserInfo{Id: userInfo.Id, Password: userInfo.Password, Name: userInfo.Name},nil
}