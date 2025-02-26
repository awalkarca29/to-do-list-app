package helper

import "to-do-list-app/entity"

type UserFormatter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FormatUser(user entity.User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	return formatter
}

type UsersFormatter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
