package seed

import (
	"roof-stack/internal/container"
	"roof-stack/internal/entity"

	log "github.com/sirupsen/logrus"

	"github.com/bxcodec/faker/v3"
)

func CreateUser() error {
	user := entity.User{}
	if err := faker.FakeData(&user); err != nil {
		log.Errorln("Error user seed", err)
	}
	db := container.DB
	return db.Create(&entity.User{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		Password: user.Password,
	}).Error
}