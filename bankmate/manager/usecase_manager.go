package manager

import "go-paybro/usecase"

type UsecaseManager interface {
	CustomerUsecase() usecase.CustomerUsecase
}

type usecaseManager struct {
	repoManager RepoManager
}

func (u *usecaseManager) CustomerUsecase() usecase.CustomerUsecase {
	return usecase.NewCustomerUsecase(u.repoManager.CustomerRepo())
}

func NewUsecaseManager(rm RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: rm,
	}
}
