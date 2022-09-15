package service

import (
	"roof-stack/internal/dto"
	"roof-stack/internal/entity"
	"roof-stack/internal/repository"

	log "github.com/sirupsen/logrus"

	"github.com/mashingan/smapping"

	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) bool
	CreateUser(user dto.RegisterDTO) entity.User
}

type authService struct {
	userRepository repository.IUserRepository
}

func NewAuthService(userRep repository.IUserRepository) *authService {
	return &authService{userRepository: userRep}
}

func (s *authService) VerifyCredential(email string, password string) interface{} {
	res := s.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (s *authService) IsDuplicateEmail(email string) bool {

	res := s.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func (s *authService) CreateUser(user dto.RegisterDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := s.userRepository.InsertUser(userToCreate)
	return res
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}