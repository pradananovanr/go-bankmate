package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"go-bankmate/model/entity"
	"time"
)

type PaymentRepo interface {
	ValidateToken(id_customer int, token string) error
	CreatePayment(id_customer int, payment *entity.PaymentRequest) (*entity.Payment, error)
	GetPayment(id_payment int) (*entity.Payment, error)
	GetAllPayment(id_customer int) ([]*entity.Payment, error)
}

type paymentRepo struct {
	db *sql.DB
}

func (p *paymentRepo) ValidateToken(id int, token string) error {
	var tokenString string

	query := "SELECT token FROM t_token WHERE id_customer = $1 AND revoked = false LIMIT 1"
	row := p.db.QueryRow(query, id)
	err := row.Scan(&tokenString)

	if err != nil {
		return err
	}

	if tokenString != token {
		return errors.New("invalid token")
	}

	return nil
}

func (p *paymentRepo) CreatePayment(id_customer int, payment *entity.PaymentRequest) (*entity.Payment, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return &entity.Payment{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		} else {
			err = tx.Commit()
			if err != nil {
				tx.Rollback()
				return
			}
		}
	}()

	var id_merchant int

	query := "SELECT id_merchant FROM m_merchant WHERE merchant_name = $1"
	row := tx.QueryRow(query, payment.Payment_Merchant)
	err = row.Scan(&id_merchant)

	if err != nil {
		return &entity.Payment{}, err
	}

	var wallet_amount float32

	query = "SELECT wallet_amount FROM t_wallet WHERE id_customer = $1"
	row = tx.QueryRow(query, id_customer)
	err = row.Scan(&wallet_amount)

	if err != nil {
		return &entity.Payment{}, err
	}

	if wallet_amount < payment.Payment_Amount {
		return &entity.Payment{}, errors.New("insufficient wallet amount")
	}

	var id_payment int
	var date_time time.Time

	query = `INSERT INTO t_payment (id_customer, id_merchant, payment_code, payment_amount, payment_description) VALUES ($1, $2, $3, $4, $5)
			RETURNING id_payment, date_time`
	row = tx.QueryRow(query, id_customer, id_merchant, payment.Payment_Code, payment.Payment_Amount, payment.Payment_Description)
	err = row.Scan(&id_payment, &date_time)

	if err != nil {
		return &entity.Payment{}, err
	}

	wallet_amount_left := wallet_amount - payment.Payment_Amount

	query = `UPDATE t_payment SET wallet_amount = $1 WHERE id_customer = $2`
	_, err = tx.Exec(query, wallet_amount_left, id_customer)

	if err != nil {
		return &entity.Payment{}, err
	}

	activity := fmt.Sprintf("customer with id %d do payment to merchant %s", id_customer, payment.Payment_Merchant)

	query = `INSERT INTO t_log (id_customer, activity) VALUES ($1, $2)`
	_, err = tx.Exec(query, id_customer, activity)
	if err != nil {
		return &entity.Payment{}, err
	}

	response := &entity.Payment{
		ID_Payment:          id_payment,
		Payment_Code:        payment.Payment_Code,
		Payment_Merchant:    payment.Payment_Merchant,
		Payment_Amount:      payment.Payment_Amount,
		Payment_Description: payment.Payment_Description,
		Date_Time:           date_time,
	}

	return response, nil
}

func (p *paymentRepo) GetPayment(id_payment int) (*entity.Payment, error) {
	var payment entity.Payment

	query := "SELECT * FROM t_payment WHERE id_payment = $1"
	row := p.db.QueryRow(query, id_payment)
	err := row.Scan(&payment)

	if err != nil {
		return &entity.Payment{}, err
	}

	return &payment, nil
}

func (p *paymentRepo) GetAllPayment(id_customer int) ([]*entity.Payment, error) {
	var payments []*entity.Payment

	query := "SELECT * FROM t_payment WHERE id_customer = $1"
	row, err := p.db.Query(query, id_customer)

	if err != nil {
		return []*entity.Payment{}, err
	}

	defer row.Close()
	for row.Next() {
		var payment entity.Payment
		if err := row.Scan(&payment.ID_Payment, &payment.Payment_Code, &payment.Payment_Merchant, &payment.Payment_Amount, &payment.Payment_Description, &payment.Date_Time); err != nil {
			return []*entity.Payment{}, err
		}
		payments = append(payments, &payment)
	}
	if err := row.Err(); err != nil {
		return []*entity.Payment{}, err
	}

	return payments, nil
}

func NewPaymentRepository(db *sql.DB) PaymentRepo {
	repo := new(paymentRepo)
	repo.db = db
	return repo
}
