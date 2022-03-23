package services

import (
	"github.com/qndaa/backend/src/repostitories"
	"github.com/qndaa/backend/src/token"
	"github.com/qndaa/backend/src/utils"
	"net/http"
	"time"
)

type UserService struct {
	userRepository *repostitories.UserRepository
}

var UserInstanceService UserService

func init() {
	UserInstanceService = UserService{
		userRepository: &repostitories.UserInstanceRepository,
	}
}

func (s *UserService) Login(w http.ResponseWriter, username, password string) {
	user, err := s.userRepository.FetchUserByUsernameAndPassword(username, password)
	if user.ID == 0 || err != nil {
		utils.BadRequest(w, "Invalid login params!")
		return
	}

	maker := token.NewJWTMaker("secret") // Trebalo bi ga cuvati u env
	t, err := maker.CreateToken(user.Username, time.Minute*30)
	w.WriteHeader(200)
	w.Write([]byte(t))

}
