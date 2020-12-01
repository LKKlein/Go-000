package service

import "week02/dao"

func FindUserById(id string) (*dao.User, error) {
	return dao.FindUserByID(id)
}
