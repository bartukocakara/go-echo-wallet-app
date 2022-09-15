package container

import (
	handler "roof-stack/internal/handler/auth"
	"roof-stack/internal/repository"
	"roof-stack/internal/service"
)

var (
	IuserRepository repository.IUserRepository = repository.NewUserRepository(DB)
	IjwtService     service.IJWTService        = service.NewJWTService()
	IauthService    service.IAuthService       = service.NewAuthService(IuserRepository)
	IauthHandler    handler.IAuthHandler       = handler.NewAuthHandler(IauthService, IjwtService)
)