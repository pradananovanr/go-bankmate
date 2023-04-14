/*
 * Author : Pradana Novan Rianto (https://github.com/pradananovanr)
 * Created on : Thu Apr 13 2023
 * Copyright : Pradana Novan Rianto © 2023. All rights reserved
 */

package repository

import (
	"database/sql"
	"fmt"
	"go-bankmate/model/entity"
	"go-bankmate/util"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type CustomerRepo interface {
	Create(newCustomer *entity.Customer) (entity.Customer, error)
	Delete(id int) error
	Login(username, password string) (string, error)
	Logout(id int) error
	InsertToken(id int, token string) error
	UpdateToken(id int) error
	AuthToken(token string) error
}

type customerRepo struct {
	db *sql.DB
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (r *customerRepo) Create(newCustomer *entity.Customer) (entity.Customer, error) {
	query := "INSERT INTO m_customer (username, password, email, phone) VALUES ($1, $2, $3, $4) RETURNING id_customer"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newCustomer.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.Customer{}, err
	}

	var customerID int
	err = r.db.QueryRow(query, newCustomer.Username, string(hashedPassword), newCustomer.Email, newCustomer.Phone).Scan(&customerID)
	if err != nil {
		log.Println(err)
		return entity.Customer{}, err
	}

	newCustomer.ID_Customer = customerID

	return *newCustomer, nil
}

func (r *customerRepo) Delete(id int) error {
	query := "DELETE FROM m_customer WHERE id_customer = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		log.Println(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Println(err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("customer with id %d not found", id)
	}

	return nil
}

func (r *customerRepo) Login(username, password string) (string, error) {

	var err error

	u := entity.CustomerLogin{}

	query := "SELECT id_customer, username, password FROM m_customer WHERE username = $1"
	row := r.db.QueryRow(query, username)
	err = row.Scan(&u.ID_Customer, &u.Username, &u.Password)

	if err != nil {
		log.Println(88)
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		log.Println(95)
		return "", err
	}

	token, err := util.GenerateToken(u.ID_Customer)

	if err != nil {
		log.Println(102)
		return "", err
	}

	err = r.InsertToken(u.ID_Customer, token)

	if err != nil {
		log.Println(109)
		return "", err
	}

	log.Println(113)
	return token, nil
}

func (r *customerRepo) Logout(id int) error {

	err := r.UpdateToken(id)

	if err != nil {
		return err
	}

	return nil
}

func (r *customerRepo) UpdateToken(id int) error {
	query := "UPDATE m_token SET revoked = $1 WHERE id_customer = $2"
	result, err := r.db.Exec(query, true, id)
	if err != nil {
		log.Println(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Println(err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("token for customer with id %d not found", id)
	}

	return nil
}

func (r *customerRepo) InsertToken(id int, token string) error {
	query := "INSERT INTO m_token (id_customer, token, revoked) VALUES ($1, $2, $3)"
	result, err := r.db.Exec(query, id, token, false)
	if err != nil {
		log.Println(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Println(err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *customerRepo) AuthToken(token string) error {
	var idToken int
	row := r.db.QueryRow("SELECT id_customer FROM m_token WHERE token = $1 AND revoked = false", token)
	err := row.Scan(idToken)

	if err != nil {
		return err
	}

	return nil
}

func NewCustomerRepository(db *sql.DB) CustomerRepo {
	repo := new(customerRepo)
	repo.db = db
	return repo
}
