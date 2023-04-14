package manager

import "go-bankmate/repository"

type RepoManager interface {
	CustomerRepo() repository.CustomerRepo
}

type repositoryManager struct {
	infraManager InfraManager
}

func (r *repositoryManager) CustomerRepo() repository.CustomerRepo {
	return repository.NewCustomerRepository(r.infraManager.DbConn())
}

func NewRepoManager(manager InfraManager) RepoManager {
	return &repositoryManager{
		infraManager: manager,
	}
}
