package getuserusecase

import (
	"go-clean-api/cmd/domain/entities/user"
	domainexceptions "go-clean-api/cmd/domain/exceptions"
	portsservice "go-clean-api/cmd/domain/services"
	domainusecases "go-clean-api/cmd/domain/use-cases"
)

type (
	getUserUseCase struct {
		UserService portsservice.UserService
	}
)

func New(us portsservice.UserService) domainusecases.GetUserUseCase {
	return &getUserUseCase{
		UserService: us,
	}
}

func (guuc *getUserUseCase) GetUser(id string) (*user.User, *domainexceptions.ApplicationException, error) {
	u, errApp, err := guuc.UserService.GetUser(id)

	return u, errApp, err
}
