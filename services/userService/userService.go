package userService

import (
	"errors"
	"luckyChess/entities"
	"luckyChess/utils"
)

type UserService struct {
	userMap map[string]entities.User
}

func New() *UserService {
	return &UserService{
		userMap: make(map[string]entities.User),
	}
}

func (g UserService) GetUser(code string) (entities.User, error) {
	hasUser, err := g.HasUser(code)

	if err != nil {
		return entities.User{}, err
	}

	if !hasUser {
		return entities.User{}, errors.New("User doesn't exists")
	}

	return g.userMap[code], nil
}

func (g UserService) GenerateUser(nickname string) (entities.User, error) {
	key := utils.RandomString(8)
	user := entities.User{
		Code:     key,
		Nickname: nickname,
	}

	g.userMap[key] = user
	return user, nil
}

func (g UserService) UpdateUser(user entities.User) error {
	hasUser, err := g.HasUser(user.Code)

	if err != nil {
		return err
	}

	if !hasUser {
		return errors.New("user doesn't exists")
	}

	g.userMap[user.Code] = user
	return nil
}

func (g UserService) DeleteUser(code string) error {
	hasUser, err := g.HasUser(code)

	if err != nil {
		return err
	}

	if !hasUser {
		return errors.New("user doesn't exists")
	}

	delete(g.userMap, code)
	return nil
}
func (g UserService) HasUser(code string) (bool, error) {

	hasCode := false
	for k := range g.userMap {
		hasCode = k == code
		if hasCode {
			break
		}
	}
	return hasCode, nil
}
