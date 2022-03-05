package handlers

import (
	"errors"
	"fmt"

	"github.com/xuanvan229/blog-core/datatransfers"
	"github.com/xuanvan229/blog-core/models"
)

func (m *module) RetrieveUser(username string) (user models.User, err error) {
	if user, err = m.db.userOrmer.GetOneByUsername(username); err != nil {
		return models.User{}, fmt.Errorf("cannot find user with username %s", username)
	}
	return
}

func (m *module) UpdateUser(id uint, user datatransfers.UserUpdate) (err error) {
	if err = m.db.userOrmer.UpdateUser(models.User{
		ID:    id,
		Email: user.Email,
		Bio:   user.Bio,
	}); err != nil {
		return errors.New("cannot update user")
	}
	return
}
