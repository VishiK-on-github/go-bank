package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*TransferRequest) error
	GetAccountByID(int) (*Account, error)
	GetAccounts() ([]*Account, error)
}

type PostgressStore struct {
	db *sql.DB
}

// used to connect to the database with creds and get a valid db object
func NewPostgressStore() (*PostgressStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgressStore{
		db: db,
	}, nil

}

// contains setup of db before services can work
func (s *PostgressStore) Init() error {
	return s.createAccountTable()
}

// used to create account table in the postgres database
func (s *PostgressStore) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS ACCOUNT (
		id SERIAL PRIMARY KEY,
		first_name varchar(50),
		last_name varchar(50),
		number serial,
		balance serial,
		created_at timestamp)`

	_, err := s.db.Exec(query)
	return err
}

// takes an account type and stores it into account table
func (s *PostgressStore) CreateAccount(account *Account) error {
	query := `INSERT INTO ACCOUNT 
	(first_name, last_name, number, balance, created_at) 
	VALUES ($1, $2, $3, $4, $5)`

	_, err := s.db.Query(query,
		account.FirstName,
		account.LastName,
		account.Number,
		account.Balance,
		account.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

// takes id of an account and deletes it from db (hard delete)
func (s *PostgressStore) DeleteAccount(id int) error {
	query := "DELETE FROM ACCOUNT WHERE ID = $1"
	_, err := s.db.Query(query, id)
	return err
}

// this method is to transfer money to an account based on its id
func (s *PostgressStore) UpdateAccount(transfer *TransferRequest) error {
	updateBalanceQuery := "UPDATE ACCOUNT SET balance = balance + $1 WHERE id = $2"
	_, err := s.db.Query(updateBalanceQuery, transfer.Amount, transfer.ToAccount)

	if err != nil {
		return err
	}

	return nil
}

// this method is used to get a user from db using its id
func (s *PostgressStore) GetAccountByID(id int) (*Account, error) {

	query := "SELECT * FROM ACCOUNT WHERE ID = $1"
	rows, err := s.db.Query(query, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account with id: %d not found", id)
}

// this method is used to get all the accounts in postgres
func (s *PostgressStore) GetAccounts() ([]*Account, error) {
	query := "SELECT * FROM ACCOUNT"

	rows, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}

	accounts := []*Account{}

	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

// utility method to iterate over rows from db
func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreatedAt)

	return account, err
}
