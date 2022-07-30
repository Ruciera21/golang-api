package domain

import (
	"database/sql"
	"goapi-nunu/errors"
	"goapi-nunu/logs"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

// Create Constructor
func NewCustomerRepositoryDB() CustomerRepositoryDB {
	//sslmode = disable selama dalam tahap development
	connStr := "postgres://postgres:%23Amanoshiru21@localhost/golang2?sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return CustomerRepositoryDB{db}
}

func (d CustomerRepositoryDB) FindByID(customerID string) (*Customer, *errors.AppError) {
	query := "select * from customers where customer_id = $1"

	// row := d.client.QueryRow(query, customerID)
	// err := row.Scan(&c.ID, &c.Name,

	var c Customer

	err := d.client.Get(&c, query, customerID)
	if err != nil {
		if err == sql.ErrNoRows {
			logs.Error("error customer data not found" + err.Error())
			return nil, errors.NewNotFoundError("customer data no found")
		} else {
			logs.Error("error scanning customer data" + err.Error())
			return nil, errors.NewUnexpectedError("unexpected database error")
		}
	}
	// untuk balikin nilai struct pakai pointers
	return &c, nil
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errors.AppError) {
	var e []Customer

	if status == "" {
		query := "select * from customers"
		err := d.client.Select(&e, query)
		if err != nil {
			// logs.Error("error fetching customer data from db: " + err.Error)
			log.Println("error query data to customer table", err.Error())
			return nil, errors.NewUnexpectedError("unexpected database error")
		}
	} else {
		if status == "active" {
			status = "1"
			query := "select * from customer where status = $1"
			err := d.client.Select(&e, query)
			if err != nil {
				log.Println("error query data to customer table", err.Error())
				return nil, errors.NewUnexpectedError("unexpected database error")
			}
		} else if status == "inactive" {
			status = "0"
			query := "select * from customer where status = $1"
			err := d.client.Select(&e, query)
			if err != nil {
				log.Println("error query data to customer table", err.Error())
				return nil, errors.NewUnexpectedError("unexpected database error")
			}
		} else {
			query := "select * from customer where status = $1"
			err := d.client.Select(&e, query)
			if err != nil {
				log.Println("error query data to customer table", err.Error())
				return nil, errors.NewUnexpectedError("unexpected database error")
			}
		}
	}
	return e, nil
}
