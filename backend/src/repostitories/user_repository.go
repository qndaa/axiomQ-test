package repostitories

import (
	"github.com/qndaa/backend/src/db"
	"github.com/qndaa/backend/src/model"
)

type UserRepository struct {
}

var UserInstanceRepository UserRepository

func init() {
	UserInstanceRepository = UserRepository{}
}

func (r *UserRepository) FetchUserByUsernameAndPassword(username, password string) (*model.User, error) {
	var user model.User
	response := db.DB.Where("username = ? AND password = ?", username, password).Find(&user)
	return &user, response.Error
}
