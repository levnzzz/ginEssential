package dto

import "github.com/levnzzz/ginEssential/model"

type  UserDto struct {
	Name string `josn:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name: user.Name,
		Telephone: user.Telephone,
	}
}