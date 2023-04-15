package usecase

import (
	"go-bankmate/model/entity"
	"go-bankmate/repository"
)

type PaymentUsecase interface {
	Create(id_customer int, payment *entity.PaymentRequest) (*entity.Payment, error)
	FindOne(id_payment int) (*entity.Payment, error)
	FindAll(id_customer int) ([]*entity.Payment, error)
}

type paymentUsecase struct {
	paymentRepo repository.PaymentRepo
}

func (u *paymentUsecase) Create(id_customer int, payment *entity.PaymentRequest) (*entity.Payment, error) {
	return u.paymentRepo.CreatePayment(id_customer, payment)
}

func (u *paymentUsecase) FindOne(id_payment int) (*entity.Payment, error) {
	return u.paymentRepo.GetPayment(id_payment)
}

func (u *paymentUsecase) FindAll(id_customer int) ([]*entity.Payment, error) {
	return u.paymentRepo.GetAllPayment(id_customer)
}

func NewPaymentUsecase(paymentRepo repository.PaymentRepo) PaymentUsecase {
	return &paymentUsecase{
		paymentRepo: paymentRepo,
	}
}
