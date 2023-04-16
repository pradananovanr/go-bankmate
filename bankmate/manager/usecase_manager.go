package manager

import "go-bankmate/usecase"

type UsecaseManager interface {
	CustomerUsecase() usecase.CustomerUsecase
	PaymentUsecase() usecase.PaymentUsecase
	DepositUsecase() usecase.DepositUsecase
}

type usecaseManager struct {
	repoManager RepoManager
}

func (u *usecaseManager) CustomerUsecase() usecase.CustomerUsecase {
	return usecase.NewCustomerUsecase(u.repoManager.CustomerRepo())
}

func (u *usecaseManager) PaymentUsecase() usecase.PaymentUsecase {
	return usecase.NewPaymentUsecase(u.repoManager.PaymentRepo())
}

func (u *usecaseManager) DepositUsecase() usecase.DepositUsecase {
	return usecase.NewDepositUsecase(u.repoManager.DepositRepo())
}

func NewUsecaseManager(rm RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: rm,
	}
}
