package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "postgres"
    dbname   = "gobankdatabase"
)

func connectDB() (*sql.DB,error)  {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// connStr := "user=postgres dbname=gobankdatabase password=postgres sslmode=disable"
	db , err := sql.Open("postgres", psqlconn)
	return db, err
}

func newPGStore() (s *PostgresStore,err error) {
	db , err := connectDB()
	if err != nil {
		return nil , err
	}
    err = db.Ping()
	if err != nil {
		if err.Error() == `pq: table "accounts" does not exist` {
			_ = s.createAccountTable()
			db, _  := connectDB()
			return &PostgresStore{
				db: db,
				}, err
		}
	}
	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init()error{
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error{
	query := `create table if not exists account (
		id serial primary key,
		first_name varchar(40),
		last_name varchar(40),
		number serial,
		balance serial,
		created_at timestamp
	)`
	_ , err := s.db.Exec(query)
	return err
}
type storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
	GetAllAccounts()([]*Account,error)
}

// CreateAccount implements storage.
func (s *PostgresStore) CreateAccount(acc *Account) error {
	query := `insert into account
	(first_name, last_name, number, balance, created_at)
	values ($1, $2, $3, $4, $5)`
	resp, err := s.db.Query(
		query,
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.Balance,
		acc.CreatedAt,
	)
	if err !=nil  {
		return err
	}
	fmt.Printf("%+v\n",resp)
	return nil
}

// UpdateAccount implements storage.
func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

// DeleteAccount implements storage.
func (s *PostgresStore) DeleteAccount(int) error {
	return nil
}

// GetAccountByID implements storage.
func (s *PostgresStore) GetAccountByID(int) (*Account, error) {
	return nil,nil
}

func (s *PostgresStore) GetAllAccounts()([]*Account,error){
	rows, err := s.db.Query(`SELECT 
	id, 
	first_name, 
	last_name, 
	"number", 
	balance, 
	created_at
	FROM account`);if err != nil{
		return nil, err
	}
	defer rows.Close()
	accounts := []*Account{}
	for rows.Next(){
		account := new(Account)
		err := rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt,
		)
		if err != nil {
            // fmt.Printf("Error scanning row: %v\n", err)
            return nil, fmt.Errorf("scan error: %v", err)
        }
		// fmt.Printf("Scanned Account: %+v\n", account)
        accounts = append(accounts, account)
	}
	if err = rows.Err(); err != nil {
		// fmt.Printf("Error during row iteration: %v\n", err)
		return nil, fmt.Errorf("rows iteration error: %v", err)
    }
	return accounts, nil
}