package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

// CreateAccount implements storage.
func (p *PostgresStore) CreateAccount(*Account) error {
	return nil
}

// UpdateAccount implements storage.
func (p *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

// DeleteAccount implements storage.
func (p *PostgresStore) DeleteAccount(int) error {
	return nil
}

// GetAccountByID implements storage.
func (p *PostgresStore) GetAccountByID(int) (*Account, error) {
	return nil,nil
}


func newPGStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=pqgotest sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error{
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error{
	query := `create table if not exists account (
		id serial primary key,
		first_name varchar(40),
		first_name varchar(40),
		number serial,
		balance serial,
		create_at timestamp
	)`

	_ , err := s.db.Exec(query)
	return err
}