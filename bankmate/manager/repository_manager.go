package manager

import "go-bankmate/repository"

type RepoManager interface {
	CustomerRepo() repository.CustomerRepo
	PaymentRepo() repository.PaymentRepo
}

type repositoryManager struct {
	infraManager InfraManager
}

func (r *repositoryManager) CustomerRepo() repository.CustomerRepo {
	return repository.NewCustomerRepository(r.infraManager.DbConn())
}

func (r *repositoryManager) PaymentRepo() repository.PaymentRepo {
	return repository.NewPaymentRepository(r.infraManager.DbConn())
}

func NewRepoManager(manager InfraManager) RepoManager {
	return &repositoryManager{
		infraManager: manager,
	}
}
